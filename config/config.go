package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Addr     string
		User     string
		Password string
		Db       string
	}
	Server struct {
		Host               string
		Port               string
		AccessTokenSecret  string
		RefreshTokenSecret string
	}
}

var c Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Sub(os.Getenv("APP_ENV")).Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return &c
}
