package models


type UserSchema struct {

	Email 	string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name 	string `json:"name" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginSchema struct {

	Email 	string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
