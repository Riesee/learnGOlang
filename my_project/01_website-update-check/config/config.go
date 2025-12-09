package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Monitor  MonitorConfig
	Redis    RedisConfig
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	LogLevel string `mapstructure:"log_level"`
	Env     string `mapstructure:"env"`
}

type MonitorConfig struct {
	URL      string `mapstructure:"url"`
	Selector string `mapstructure:"selector"`
	Interval time.Duration `mapstructure:"interval"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

