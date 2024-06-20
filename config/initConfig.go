package config

import "github.com/spf13/viper"

func InitConfig() error {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	return err
}
