package config

import (
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.SetDefault("BaseUrl", "http://simstim.de/")
	viper.SetDefault("Token", "foobar")
	viper.SetDefault("Port", "8080")
	viper.SetDefault("IP", "localhost")
	viper.SetDefault("urlPrefix", "s")

	viper.SetEnvPrefix("LS")
	viper.BindEnv("BaseUrl")
	viper.BindEnv("Token")
	viper.BindEnv("Port")
	viper.BindEnv("IP")
	viper.BindEnv("urlPrefix")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := viper.SafeWriteConfig()
			if err != nil {
				println(err.Error())
			}
		} else {
			// Config file was found but another error was produced
			println(err.Error())
		}
	}
}
