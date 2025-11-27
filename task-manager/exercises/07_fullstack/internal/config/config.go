package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
}

type AppConfig struct {
	Name    string
	Version string
	Env     string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret     string
	ExpireHour int
}

type RedisConfig struct {
	Host string
	Port string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config okunamadı: %w", err))
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("config okunamadı: %w", err))
	}

	return &cfg
}
