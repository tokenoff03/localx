package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DB     DatabaseConfig `yaml:"db"`
	Server ServerConfig   `yaml:"http"`
}

type DatabaseConfig struct {
	URI string `yaml:"url"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
