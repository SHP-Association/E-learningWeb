package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SHP-Association/E-learningWeb/backend-go/internal/api"
	"github.com/SHP-Association/E-learningWeb/backend-go/internal/config"
	"github.com/SHP-Association/E-learningWeb/backend-go/internal/store/postgres"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("🚀 Starting E-Learning Backend (Go) on port %s", cfg.Server.Port)
	log.Printf("📝 Environment: %s", cfg.Server.Env)

	// Connect to database
	dsn := cfg.Database.GetDSN()
	store, err := postgres.New(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("✅ Database connection established")

	// Create API instance
	apiInstance := api.New(api.Opts{
		Store:         store,
		Config:        cfg,
		SessionSecret: os.Getenv("GO_JWT_SECRET"),
	})

	// Setup HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      apiInstance.Routes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server
	log.Printf("🌐 Server listening on http://localhost:%s", cfg.Server.Port)
	log.Printf("📚 API endpoints available at http://localhost:%s/api/", cfg.Server.Port)
	log.Printf("💚 Health check: http://localhost:%s/api/health", cfg.Server.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
