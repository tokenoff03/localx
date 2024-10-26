package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DB     DatabaseConfig `yaml:"db"`
	Server ServerConfig   `yaml:"http"` // Секция "http" в YAML соответствует полю Server
	AWS    AWSConfig      `yaml:"aws"`
}

type DatabaseConfig struct {
	URL string `yaml:"url"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type AWSConfig struct {
	S3 S3Config `yaml:"s3"`
}
type S3Config struct {
	Bucket string `yaml:"bucket"`
	Region string `yaml:"region"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
