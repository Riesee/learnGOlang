package main

import "fmt"

type Config struct {
	databaseURL string
	redisURL    string
}

type Database struct {
	URL string
}

type Cache struct {
	URL string
}

type UserRepository struct {
	DB *Database
}


type UserService struct {
	repository *UserRepository
	cache      *Cache
}

func NewConfig() *Config{
	return &Config{
		databaseURL: "postgres://localhost:5432",
		redisURL:    "redis://localhost:6379",
	}
}

func NewDatabase(cfg *Config) *Database {
	fmt.Println("Database bağlanıyor...", cfg.databaseURL)
	return &Database{
		URL: cfg.databaseURL,
	}
}

func NewCache(cfg *Config) *Cache {
	fmt.Println("Cache bağlanıyor...", cfg.redisURL)
	return &Cache{
		URL: cfg.redisURL,
	}
}

func NewUserRepository(database *Database) *UserRepository {
	fmt.Println("User repository oluşturuluyor...")
	return &UserRepository{
		DB: database,
	}
}

func NewUserService(repository *UserRepository, cache *Cache) *UserService {
	fmt.Println("User service oluşturuluyor...")
	return &UserService{
		repository: repository,
		cache:      cache,
	}
}

func main(){
	service := InitializeUserService()

	fmt.Printf("\nService Hazır: %v\n", service)
}