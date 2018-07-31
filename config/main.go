package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	config *viper.Viper
	Rules  map[string]interface{}
}

func New(path string) (*Config, error) {
	config := viper.New()

	config.SetConfigName("linteum")
	config.SetConfigFile(path)
	config.AddConfigPath(".")
	config.SetConfigType("yaml")

	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}

	result := &Config{
		config: config,
		Rules:  config.GetStringMap("rules"),
	}

	return result, nil
}
