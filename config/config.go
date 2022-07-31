package config

import (
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("LS")
	viper.SetDefault("BaseUrl", "http://simstim.de/s/")
	viper.SetDefault("Token", "foobar")
	viper.BindEnv("BaseUrl")
	viper.BindEnv("Token")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := viper.SafeWriteConfigAs("./config.toml")
			if err != nil {
				println(err.Error())
			}
		}
	} else {
		// Config file was found but another error was produced
		println(err.Error())
	}
}
