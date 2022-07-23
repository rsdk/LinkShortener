package config

import "github.com/spf13/viper"

func Init() {
	viper.SetEnvPrefix("LS")
	viper.SetDefault("BaseUrl", "https://simstim.de/")
	viper.SetDefault("Token", "foobar")
	viper.BindEnv("BaseUrl")
	viper.BindEnv("Token")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.WriteConfig()
		} else {
			// Config file was found but another error was produced
			println(err.Error())
		}
	}
}
