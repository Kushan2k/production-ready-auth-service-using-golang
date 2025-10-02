package main

import (
	"fmt"
	"github/go_auth_api/internal/config"
	"github/go_auth_api/internal/db"
)

func main() {
	// Entry point of the application

	fmt.Println("Hello, World!")

	cfg, err := config.LoadConfig()

	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	databaseService := &db.DataBaseService{}
	_, err = databaseService.Connect(cfg)

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	fmt.Println("Connected to database successfully!")

	fmt.Print(cfg.JWT_Secret)

}
