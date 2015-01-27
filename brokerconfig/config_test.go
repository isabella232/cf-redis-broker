package brokerconfig_test

import (
	"os"
	"path"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/cf-redis-broker/brokerconfig"
)

var _ = Describe("parsing the broker config file", func() {

	Describe("ParseConfig", func() {
		var (
			config         brokerconfig.Config
			configPath     string
			parseConfigErr error
		)

		BeforeEach(func() {
			configPath = "test_config.yml"
		})

		JustBeforeEach(func() {
			path, err := filepath.Abs(path.Join("assets", configPath))
			Ω(err).ToNot(HaveOccurred())
			config, parseConfigErr = brokerconfig.ParseConfig(path)
		})

		Context("when the configuration is valid", func() {

			var dirs []string

			BeforeEach(func() {
				dirs = []string{"/tmp/to/redis", "/tmp/redis/data/directory", "/tmp/redis/log/directory"}
				for _, dir := range dirs {
					err := os.MkdirAll(dir, 0755)
					Ω(err).ShouldNot(HaveOccurred())
				}
				_, err := os.Create("/tmp/to/redis/config.conf")
				Ω(err).ShouldNot(HaveOccurred())
			})

			AfterEach(func() {
				for _, dir := range dirs {
					err := os.RemoveAll(dir)
					Ω(err).ShouldNot(HaveOccurred())
				}
			})

			It("does not error", func() {
				Ω(parseConfigErr).NotTo(HaveOccurred())
			})

			It("loads host", func() {
				Ω(config.RedisConfiguration.Host).To(Equal("example.com"))
			})

			It("loads port", func() {
				Ω(config.Port).Should(Equal("3000"))
			})

			It("loads path to default brokerconfig.conf", func() {
				Ω(config.RedisConfiguration.DefaultConfigPath).To(Equal("/tmp/to/redis/config.conf"))
			})

			It("loads the auth credendials", func() {
				Ω(config.AuthConfiguration.Username).To(Equal("admin"))
				Ω(config.AuthConfiguration.Password).To(Equal("secret"))
			})

			It("loads the start Redis timeout", func() {
				Ω(config.RedisConfiguration.StartRedisTimeoutSeconds).To(Equal(3))
			})

			It("loads process check interval", func() {
				Ω(config.RedisConfiguration.ProcessCheckIntervalSeconds).To(Equal(5))
			})

			It("loads instance data directory", func() {
				Ω(config.RedisConfiguration.InstanceDataDirectory).To(Equal("/tmp/redis/data/directory"))
			})

			It("loads instance log directory", func() {
				Ω(config.RedisConfiguration.InstanceLogDirectory).To(Equal("/tmp/redis/log/directory"))
			})

			It("loads service instance limit", func() {
				Ω(config.RedisConfiguration.ServiceInstanceLimit).To(Equal(3))
			})

			It("loads the auth credendials", func() {
				Ω(config.AuthConfiguration.Username).To(Equal("admin"))
				Ω(config.AuthConfiguration.Password).To(Equal("secret"))
			})

			It("loads the backup configuration details", func() {
				Ω(config.RedisConfiguration.BackupConfiguration.EndpointUrl).Should(Equal("http://s3url.com"))
				Ω(config.RedisConfiguration.BackupConfiguration.BucketName).Should(Equal("redis-backups"))
				Ω(config.RedisConfiguration.BackupConfiguration.AccessKeyId).Should(Equal("ABCDEABCDEABCDEABCDE"))
				Ω(config.RedisConfiguration.BackupConfiguration.SecretAccessKey).Should(Equal("ABCDEABCDEABCDEABCDEABCDEABCDEABCDEABCDE"))
				Ω(config.RedisConfiguration.BackupConfiguration.SecurityToken).Should(Equal("ABZEVEIEIEIE"))
				Ω(config.RedisConfiguration.BackupConfiguration.Path).Should(Equal("/home"))
				Ω(config.RedisConfiguration.BackupConfiguration.S3Region).Should(Equal("france"))

				Ω(config.RedisConfiguration.BackupConfiguration.BGSaveTimeoutSeconds).To(Equal(600))
			})

			It("backup is enabled", func() {
				Ω(config.RedisConfiguration.BackupConfiguration.Enabled()).To(BeTrue())
			})

			It("loads the monit exectuable path", func() {
				Ω(config.MonitExecutablePath).Should(Equal("/some/path/to/monit"))
			})

			It("loads the redis-server exectuable path", func() {
				Ω(config.RedisServerExecutablePath).Should(Equal("/some/path/to/redis-server"))
			})
		})

		Context("when the configuration is invalid", func() {

			BeforeEach(func() {
				configPath = "test_config.yml-invalid"
			})

			It("returns an error", func() {
				Ω(parseConfigErr).Should(MatchError(ContainSubstring("not found")))
			})
		})

		Context("when the backup configuration is empty", func() {
			BeforeEach(func() {
				configPath = "test_config.yml-no-backup"
			})

			It("backup is disabled", func() {
				Ω(config.RedisConfiguration.BackupConfiguration.Enabled()).To(BeFalse())
			})
		})

		Context("when the backup configuration has minimum to start a backup", func() {
			BeforeEach(func() {
				configPath = "test_config.yml-backup-minimum"
			})

			It("backup is enabled", func() {
				Ω(config.RedisConfiguration.BackupConfiguration.Enabled()).To(BeTrue())
			})
		})

		Describe("dedicated nodes", func() {
			It("loads the dedicated node ips", func() {
				Ω(len(config.RedisConfiguration.Dedicated.Nodes)).Should(Equal(3))
				Ω(config.RedisConfiguration.Dedicated.Nodes).Should(ContainElement("10.0.0.1"))
				Ω(config.RedisConfiguration.Dedicated.Nodes).Should(ContainElement("10.0.0.2"))
				Ω(config.RedisConfiguration.Dedicated.Nodes).Should(ContainElement("10.0.0.3"))
			})

			It("sets the correct port", func() {
				Ω(config.RedisConfiguration.Dedicated.Port).Should(Equal(6379))
			})

			It("sets the path to the statefile", func() {
				Ω(config.RedisConfiguration.Dedicated.StatefilePath).Should(Equal("/tmp/redis-config-dir/statefile.json"))
			})
		})
	})

	Describe("ValidateConfig", func() {
		var validFile string
		var validDir string
		var err error
		var config brokerconfig.ServiceConfiguration

		BeforeEach(func() {
			validFile, err = filepath.Abs(path.Join("assets", "test_config.yml"))
			Ω(err).ToNot(HaveOccurred())

			validDir, err = filepath.Abs(path.Join("assets"))
			Ω(err).ToNot(HaveOccurred())

			config = brokerconfig.ServiceConfiguration{
				DefaultConfigPath:     validFile,
				InstanceDataDirectory: validDir,
				InstanceLogDirectory:  validDir,
			}
		})

		Describe("DefaultRedisConfPath", func() {
			Context("When the default redis conf path points to an existing file", func() {
				It("does not return an error", func() {
					err := brokerconfig.ValidateConfig(config)
					Ω(err).ToNot(HaveOccurred())
				})
			})

			Context("When the default redis conf path points to a non-existent file", func() {
				It("returns an error", func() {
					config.DefaultConfigPath = "/a/non-existent/path"
					err := brokerconfig.ValidateConfig(config)
					Ω(err).To(HaveOccurred())
					Ω(err.Error()).To(Equal("File '/a/non-existent/path' (RedisConfig.DefaultRedisConfPath) not found"))
				})
			})
		})

		Describe("InstanceDataDirectory", func() {
			Context("When the instance data directory path points to an existing directory", func() {
				It("does not return an error", func() {
					err := brokerconfig.ValidateConfig(config)
					Ω(err).ToNot(HaveOccurred())
				})
			})

			Context("When the instance data directory points to a non-existent directory", func() {
				It("returns an error", func() {
					config.InstanceDataDirectory = "/a/non-existent/path"
					err := brokerconfig.ValidateConfig(config)
					Ω(err).To(HaveOccurred())
					Ω(err.Error()).To(Equal("File '/a/non-existent/path' (RedisConfig.InstanceDataDirectory) not found"))
				})
			})
		})

		Describe("InstanceLogDirectory", func() {
			Context("When the instance log directory path points to an existing directory", func() {
				It("does not return an error", func() {
					err := brokerconfig.ValidateConfig(config)
					Ω(err).ToNot(HaveOccurred())
				})
			})

			Context("When the instance log directory points to a non-existent directory", func() {
				It("returns an error", func() {
					config.InstanceLogDirectory = "/a/non-existent/path"
					err := brokerconfig.ValidateConfig(config)
					Ω(err).To(HaveOccurred())
					Ω(err.Error()).To(Equal("File '/a/non-existent/path' (RedisConfig.InstanceLogDirectory) not found"))
				})
			})
		})
	})
})
