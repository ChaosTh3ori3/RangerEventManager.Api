package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(filePath string) (Config, error) {
	var config Config

	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
