package services

import (
	"bytes"
	"github/go_auth_api/internal/config"
	"github/go_auth_api/internal/models"
	"github/go_auth_api/internal/utils"
	"html/template"

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
		Name: 	 new_user.Name,
	}

	if err := s.DB.Create(user).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error creating user")
	}

	mailer:=utils.GetMailer()
	msg:=utils.GetMessage()

	t,err:=template.ParseFiles("./internal/emails/register_email.html")

	if err ==nil {
		var data bytes.Buffer

		//Customize what you are passing to the email template
		//For example, you can pass the verification code or user name
		//Here, we are passing the user model with Name and Email fields
		//You can modify the template to use these fields as needed
		if err:=t.Execute(&data,user);err==nil{
			msg.SetBody("text/html",data.String())
		}
		msg.SetHeader("From", config.Envs.MAIL_USERNAME)
		msg.SetHeader("To", new_user.Email)
		msg.SetHeader("Subject", "Welcome to Our Service")

		go mailer.DialAndSend(msg)
	}

	

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
		"status":fiber.StatusCreated,
	})
}

func (s *AuthService) LoginUser(c *fiber.Ctx) error {
	

	login_user:=new(models.LoginSchema)

	if err:=c.BodyParser(login_user);err!=nil{
		return fiber.NewError(fiber.StatusBadRequest,err.Error())
	}

	var user models.User
	if err:=s.DB.Where("email = ?", login_user.Email).First(&user).Error;err!=nil{
		return fiber.NewError(fiber.StatusUnauthorized,"Invalid email or password")
	}

	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(login_user.Password));err!=nil{
		return fiber.NewError(fiber.StatusUnauthorized,"Invalid email or password")
	}

	token,err:=utils.GenerateJWT(&user,s.cfg.JWT_SECRET)

	if err!=nil {
		return fiber.NewError(fiber.StatusInternalServerError,"Error generating token")
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Login successful",
		"token":token,
		"status":fiber.StatusOK,
	})

}
