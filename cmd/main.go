// @title Fiber Authentication API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email kushangayantha001@gmail.com
// @license.name MIT
// @license.url http://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api/v1/auth
package main

import (
	"fmt"
	"github/go_auth_api/internal/config"
	"github/go_auth_api/internal/db"
	"github/go_auth_api/internal/middlewares"
	"github/go_auth_api/internal/models"
	"github/go_auth_api/internal/router"

	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/swagger"

	_ "github/go_auth_api/docs"

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
	error_handler := middlewares.NewErrorHandler(cfg)

	engine:=html.New("./internal/emails", ".html")

	app := fiber.New(
		fiber.Config{
			Prefork:       false,
			CaseSensitive: true,
			StrictRouting: true,
			ErrorHandler:  error_handler.HandleError,
			Views:         engine,
			
		},
	)
	app.Get("/docs/*", swagger.HandlerDefault)


	api := app.Group("/api/v1/auth")

	auth_routes := router.NewAuthRouter(api, gormDB, cfg)
	auth_routes.SetupRoutes()

	// if err := app.Listen(fmt.Sprintf(":%s", cfg.SERVER_PORT)); err != nil {
	// 	fmt.Println("Error starting server:", err)
	// 	return
	// }

}
