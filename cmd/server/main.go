package main

import (
	"fmt"
	"lexcodex/cmd/auth"
	"lexcodex/config"
	"lexcodex/internal/models"
	"lexcodex/internal/routes"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title           LexCodex API
// @version         1.0
// @description     Personal expense tracking API with 50/30/20 budget management

// @contact.name   Juan Vargas
// @contact.url    https://github.com/jpvargasdev/LexCodex

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Firebase ID Token (Bearer token)

// @externalDocs.description  OpenAPI

func main() {
	config.Load()

	models.InitializeDatabase()
	defer models.CloseDatabase()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		fmt.Println("\nShutting down server...")
		models.CloseDatabase()
		os.Exit(0)
	}()

	if err := models.CreateTables(); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Seed the database with initial categories
	if err := models.SeedCategories(); err != nil {
		log.Fatalf("Failed to seed categories: %v", err)
	}

	// Init Firebase
	err := auth.InitFirebase()
	if err != nil {
		fmt.Printf("Failed to initialize Firebase: %v", err)
	}

	router := routes.SetupRouter()

	port := config.GetServerPort()

	fmt.Printf("LexCodex server is running on port %s...\n", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting LexCodex server: %v", err)
	}
}
