package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port             string `mapstructure:"PORT"`
	UserSvcUrl       string `mapstructure:"USER_SVC_URL"`
	ProductSvcUrl    string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl      string `mapstructure:"ORDER_SVC_URL"`
	AdminSvcUrl      string `mapstructure:"ADMIN_SVC_URL"`
	CartSvcUrl       string `mapstructure:"CART_SVC_URL"`
	RoomSvcUrl       string `mapstructure:"ROOM_SVC_URL"`
	ConnectionSvcUrl string `mapstructure:"CONNECTION_SVC_URL"`
}

var envs = []string{
	"PORT", "USER_SVC_URL", "PRODUCT_SVC_URL", "ORDER_SVC_URL", "ADMIN_SVC_URL", "CART_SVC_URL", "ROOM_SVC_URL", "CONNECTION_SVC_URL",
}

func LoadConfig() (Config, error) {
	var config Config

	// Automatically look for environment variables with the same name
	viper.AutomaticEnv()

	// Optionally read from a .env file if it exists (useful for local development and Docker)
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	// Bind each environment variable
	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	fmt.Println(config.AdminSvcUrl)

	return config, nil
}
