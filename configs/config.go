package configs

import "github.com/spf13/viper"

func SetViper() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.AddConfigPath("../../configs")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
