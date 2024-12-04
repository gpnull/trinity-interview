package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type HTTPServerConfig struct {
	IP   string
	Port string
}

type DataBaseConfig struct {
	URL string
}

type Config struct {
	HTTPServer     HTTPServerConfig
	DataBaseConfig DataBaseConfig
}

func LoadConfig(configFilePath string) (*Config, error) {
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
