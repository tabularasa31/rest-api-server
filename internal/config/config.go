package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"rest-api-server/pkg/logger"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" enc-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"tcp"`
		BindIp string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
}

var cfg *Config
var once sync.Once

func GetConfig(logger *logger.Logger) *Config {
	once.Do(func() {
		logger.Info("getting config")
		cfg = &Config{}
		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			info, _ := cleanenv.GetDescription(cfg, nil)
			logger.Info(info)
			logger.Fatalln(err)
		}
	})
	return cfg
}
