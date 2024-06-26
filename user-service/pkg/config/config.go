package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost      string `mapstructure:"DB_HOST"`
	DBName      string `mapstructure:"DB_NAME"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	Port        string `mapstructure:"PORT"`
	MongoDBHost string `mapstructure:"MONGO_DB_HOST"`

	MongoDBDatabase       string `mapstructure:"MONGO_DB_DATABASE"`
	MongoDBPort           string `mapstructure:"MONGO_DB_PORT"`
	AdminSvcUrl           string `mapstructure:"ADMIN_SVC_URL"`
	ChatSvcUrl            string `mapstructure:"CHAT_SVC_URL"`
	ConnectionSvcUrl      string `mapstructure:"CONNECTION_SVC_URL"`
	Email                 string `mapstructure:"email"`
	Password              string `mapstructure:"password"`
	KafkaTopic            string `mapstructure:"KAFKA_TOPIC"`
	KafkaBrokers          string `mapstructure:"KAFKA_BROKERS"`
	AWS_REGION            string `mapstructure:"AWS_REGION`
	AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID`
	AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "PORT", "ADMIN_SVC_URL", "CHAT_SVC_URL", "CONNECTION_SVC_URL", "MONGO_DB_HOST", "MongoDBPort", "MONGO_DB_DATABASE",
	"AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY",
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
