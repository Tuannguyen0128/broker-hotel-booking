package configs

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
	"sync"
)

var (
	lock = &sync.Mutex{}
	// ConfigInstance ...
	ConfigInstance *AppConfig
)

// AppConfig holds the application's configuration parameters.
type AppConfig struct {
	RepoServer       string   `mapstructure:"repo_server"`
	BrokerServer     string   `mapstructure:"broker_server"`
	KafkaServer      string   `mapstructure:"kafka_server"`
	KafkaBrokerTopic string   `mapstructure:"kafka_broker-topic"`
	KafkaRepoTopic   string   `mapstructure:"kafka_repo-topic"`
	MongoDB          DBConfig `mapstructure:"mongo-db"`
}

type DBConfig struct {
	Host   string `mapstructure:"host"`
	Schema string `mapstructure:"schema"`
}

// Load ...
func Load(cfg string) *AppConfig {
	if ConfigInstance == nil {
		cfg := strings.TrimSpace(cfg)
		if "" == cfg {
			log.WithError(nil).Error("No configuration file specified")
			os.Exit(1)
		}

		lock.Lock()
		defer lock.Unlock()

		log.Println("Creating single instance of the AppConfig")

		var cfgOptions AppConfig
		viper.SetConfigFile(cfg)
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
		viper.AutomaticEnv() // read in environment variables that match
		// If a configs file is found, read it in.
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Error("Could not read configs file")
			os.Exit(1)
		}
		//load into conf format
		if err := viper.Unmarshal(&cfgOptions); err != nil {
			log.WithError(err).Error("Could not unmarshal configs into conf var")
			os.Exit(1)
		}
		ConfigInstance = &cfgOptions
	}

	return ConfigInstance
}
