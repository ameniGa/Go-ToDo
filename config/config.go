package config

import (
	"github.com/jinzhu/configor"
	"path"
	"runtime"
)

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

// GetConfig parses the config file to the struct Config
func GetConfig() (*Config, error) {
	config := new(Config)
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "../config.yaml")

	err := configor.Load(config, filepath)
	return config, err
}
