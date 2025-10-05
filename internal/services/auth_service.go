package services

import (
	"github/go_auth_api/internal/config"
	"github/go_auth_api/internal/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

	new_user:=new(models.UserSchema)

	if err:=c.BodyParser(new_user);err!=nil{
		return fiber.NewError(fiber.StatusBadRequest,err.Error())
	}

	if new_user.Password!=new_user.ConfirmPassword {
		return fiber.NewError(fiber.StatusBadRequest,"Password and Confirm Password do not match")
	}

	pw_hash, err := bcrypt.GenerateFromPassword([]byte(new_user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error hashing password")
	}

	user := &models.User{
		Email:    new_user.Email,
		Password: string(pw_hash),
	}

	if err := s.DB.Create(user).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error creating user")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user":user,
		"status":fiber.StatusCreated,
	})
}

func (s *AuthService) LoginUser(c *fiber.Ctx) error {
	// Implement user login logic here
	return c.SendString("User logged in")
}
