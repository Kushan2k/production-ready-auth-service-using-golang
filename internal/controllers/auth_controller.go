package controllers

import (
	"github/go_auth_api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	return c.SendString("User registered")
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	return c.SendString("User logged in")
}
