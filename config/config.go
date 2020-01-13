package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Grpc struct {
			Host string
		}
		Http struct {
			Host string
		}
	}
	Database struct {
		Type       string
		Uri        string
		Name       string
		Collection string
	}
}

func setViper() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	return err
}

func GetConfig() (*Config, error) {
	err := setViper()
	if err != nil {
		return nil, err
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
