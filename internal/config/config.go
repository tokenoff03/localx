package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DB       DatabaseConfig `yaml:"db"`
	Server   ServerConfig   `yaml:"http"`
	AWS      AWSConfig      `yaml:"aws"`
	Supabase SupabaseConfig `yaml:"supabase"`
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

type SupabaseConfig struct {
	StorageURL string `yaml:"storageURL"`
	Bucket     string `yaml:"bucket"`
	AuthToken  string `yaml:"authToken"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
