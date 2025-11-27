package config

import (
	"strings"

	"github.com/spf13/viper"
)

// DERS 2: Viper ile Configuration
// ================================
//
// Viper özellikleri:
// - YAML, JSON, TOML, ENV dosyaları okur
// - Environment variable'ları otomatik bind eder
// - Default değerler tanımlayabilirsin
// - Hot reload destekler
//
// Öncelik sırası (yüksekten düşüğe):
// 1. Environment variables
// 2. Config dosyası
// 3. Default değerler

// Config ana konfigürasyon struct'ı
type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
}

type AppConfig struct {
	Name string
	Env  string // development, production
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
	Host     string
	Port     string
	Password string
	DB       int
}

// Load konfigürasyonu yükler
func Load() (*Config, error) {
	// Default değerler
	viper.SetDefault("app.name", "task-manager")
	viper.SetDefault("app.env", "development")
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("jwt.expireHour", 24)
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.db", 0)

	// Config dosyası ayarları
	viper.SetConfigName("config")    // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")         // Çalışma dizini
	viper.AddConfigPath("./config")  // config/ klasörü

	// Environment variable'ları oku
	// DATABASE_HOST -> database.host olur
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Config dosyasını oku (yoksa hata verme, env var'lar yeterli)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	// Struct'a unmarshal et
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// GetDSN PostgreSQL connection string döndürür
func (c *DatabaseConfig) GetDSN() string {
	return "host=" + c.Host +
		" user=" + c.User +
		" password=" + c.Password +
		" dbname=" + c.DBName +
		" port=" + c.Port +
		" sslmode=" + c.SSLMode
}
