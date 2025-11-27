package database

import (
	"fullstack/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.DatabaseConfig) (*gorm.DB, error) {

	postgresInfo := "host=" + cfg.Host + " user=" + cfg.User + " password=" + cfg.Password + " dbname=" + cfg.DBName + " port=" + cfg.Port + " sslmode=" + cfg.SSLMode
	db, err := gorm.Open(postgres.Open(postgresInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
