package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return Config{}, fmt.Errorf("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return Config{}, fmt.Errorf("cannot read config: %s", configPath)
	}

	return cfg, nil
}
