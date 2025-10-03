package router

import (
	"github/go_auth_api/internal/config"
	"github/go_auth_api/internal/controllers"
	"github/go_auth_api/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthRouter struct {
	Group *fiber.Router
	DB    *gorm.DB
	cfg   *config.Config
}

func NewAuthRouter(group *fiber.Router, db *gorm.DB, cfg *config.Config) *AuthRouter {
	return &AuthRouter{
		Group: group,
		DB:    db,
		cfg:   cfg,
	}
}

func (ar *AuthRouter) SetupRoutes() {
	service := services.NewAuthService(ar.DB, ar.cfg)
	_ = controllers.NewAuthController(service)

}
