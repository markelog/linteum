package config

import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

type Config struct {
	Rules map[string]string `json:"rules"`
}

var conf Config

func New(path string) (Config, error) {
	var (
		configuration = config.NewConfig()
		err           = configuration.Load(file.NewSource(
			file.WithPath(path),
		))
	)

	if err != nil {
		return conf, err
	}

	configuration.Scan(&conf)

	return conf, nil
}
