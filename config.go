package main

import (
	"feedback/logger"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Host   string
	Port   int
	User   string
	Pass   string
	DbName string
}

func loadConfig() DbConfig {
	viper.SetDefault("db-host", "localhost")
	viper.SetDefault("db-port", 5432)
	viper.SetDefault("db-name", "postgres")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger().Errorw("Unable to read config. Please ensure that config.yaml exists.",
			"error", err)
	}

	return DbConfig{
		Host:   viper.GetString("db-host"),
		Port:   viper.GetInt("db-port"),
		User:   viper.GetString("db-user"),
		Pass:   viper.GetString("db-pass"),
		DbName: viper.GetString("db-name"),
	}
}
