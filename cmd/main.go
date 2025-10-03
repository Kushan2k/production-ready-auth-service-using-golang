package main

import (
	"fmt"
	"github/go_auth_api/internal/config"
	"github/go_auth_api/internal/db"
	"github/go_auth_api/internal/models"
	"github/go_auth_api/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Entry point of the application

	fmt.Println("Hello, World!")

	cfg, err := config.LoadConfig()

	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	databaseService := db.NewDatabaseService(cfg)

	gormDB, err := databaseService.Connect()

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	gormDB.AutoMigrate(&models.User{})

	fmt.Println("Connected to database successfully!")

	app := fiber.New(
		fiber.Config{
			Prefork:       false,
			CaseSensitive: true,
			StrictRouting: true,
		},
	)

	api := app.Group("/api")

	auth_routes := router.NewAuthRouter(&api, gormDB, cfg)
	auth_routes.SetupRoutes()

	if err := app.Listen(fmt.Sprintf(":%s", cfg.Server_Port)); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

}
