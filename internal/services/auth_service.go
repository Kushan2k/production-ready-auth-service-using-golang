package services

import (
	"github/go_auth_api/internal/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthService struct {
	DB  *gorm.DB
	cfg *config.Config
}

func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{
		DB:  db,
		cfg: cfg,
	}
}

func (s *AuthService) RegisterUser(c *fiber.Ctx) error {
	// Implement user registration logic here
	return c.SendString("User registered")
}

func (s *AuthService) LoginUser(c *fiber.Ctx) error {
	// Implement user login logic here
	return c.SendString("User logged in")
}
