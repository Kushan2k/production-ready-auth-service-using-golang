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

// Register User
// @Summary      Register a new user
// @Description  Register a new user with provided details
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body models.UserSchema true "User registration payload" 
// @Router       /register [post]
func (ac *AuthController) Register(c *fiber.Ctx) error {
	
	return ac.service.RegisterUser(c)


}	

// Login user
// @Summary      user Login
// @Description  Login an existing user with provided credentials
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body models.LoginSchema true "User login payload" 
// @Router       /login [post]
func (ac *AuthController) Login(c *fiber.Ctx) error {
	return ac.service.LoginUser(c)
}
