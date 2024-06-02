package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost       string `mapstructure:"DB_HOST"`
	DBName       string `mapstructure:"DB_NAME"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	Port         string `mapstructure:"PORT"`
	MongoDBHost  string `mapstructure:"MONGO_DB_HOST"`
	MongoDBPort  string `mapstructure:"MONGO_DB_PORT"`
	AdminSvcUrl  string `mapstructure:"ADMIN_SVC_URL"`
	Email        string `mapstructure:"email"`
	Password     string `mapstructure:"password"`
	KafkaTopic   string `mapstructure:"KAFKA_TOPIC"`
	KafkaBrokers string `mapstructure:"KAFKA_BROKERS"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "PORT", "ADMIN_SVC_URL",
	"KAFKA_TOPIC", "KAFKA_BROKERS",
}

func LoadConfig() (Config, error) {
	var config Config
	viper.AutomaticEnv()
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
