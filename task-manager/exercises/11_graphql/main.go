package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"graphql-demo/graph/generated"
	"graphql-demo/graph/resolverfn"
	"graphql-demo/internal/config"
	"graphql-demo/internal/database"
	"graphql-demo/internal/middleware"
)

const defaultPort = "8080"

func main() {
	// Config yükle
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Config yüklenemedi: %v", err)
	}

	// Veritabanına bağlan
	db, err := database.Connect(&cfg.Database)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}
	log.Println("✅ Veritabanına bağlandı")

	// Resolver oluştur (DB ile)
	resolver := resolverfn.NewResolver(db)

	// GraphQL server oluştur
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolver},
		),
	)

	// Gin router oluştur
	r := gin.Default()

	// Middleware
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
	protected.Use(middleware.GinContextToContext())
	{
		protected.POST("/query", gin.WrapH(srv))
	}
	// Routes
	r.GET("/", gin.WrapH( playground.Handler("GraphQL Playground", "/query")))
	
	log.Println("===========================================")
	log.Printf("🚀 GraphQL Playground: http://localhost:%s/", defaultPort)
	log.Printf("📡 GraphQL Endpoint:   http://localhost:%s/query", defaultPort)
	log.Println("===========================================")
	r.Run(":" + defaultPort)
}
