package resolverfn

import "gorm.io/gorm"

// Resolver - dependency injection için
// DB bağlantısını tutar, tüm resolver'lar buna erişir
type Resolver struct {
	DB *gorm.DB
}

// NewResolver - yeni resolver oluşturur
func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{DB: db}
}
