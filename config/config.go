package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleconns int
		MaxOpenconns int
	}
}

var AppConfig *Config

func InitConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error unmarshalling config, %v", err)
	}

	initDB()
}
