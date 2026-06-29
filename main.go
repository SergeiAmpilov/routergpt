package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"routergpt/internal/config"
	"routergpt/internal/domain/completion/handler"
	modelshandler "routergpt/internal/domain/models/handler"
	modelsservice "routergpt/internal/domain/models/service"
	modelsrepo "routergpt/internal/domain/models/repository"
	"routergpt/internal/transport/rest"
)

func main() {
	cfg := config.New().GetConfig()

	app := fiber.New()

	// Initialize completion handler
	completionHandler := handler.New()

	// Initialize models repository
	modelsRepo, err := modelsrepo.NewPostgreSQLRepository(cfg.ConnString)
	if err != nil {
		log.Fatalf("Failed to initialize models repository: %v", err)
	}

	// Initialize models service
	modelsService := modelsservice.NewModelsService(modelsRepo)

	// Initialize models handler
	modelsHandler := modelshandler.New(modelsService)

	rest.SetupRoutes(app, completionHandler, modelsHandler)

	log.Fatal(app.Listen(":" + cfg.Port))
}