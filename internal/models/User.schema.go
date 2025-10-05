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


type VerifyOtpSchema struct {

	Email 	string `json:"email" validate:"required,email"`
	Otp 	string `json:"otp" validate:"required,min=6,max=6"`
}

type ResendOtpSchema struct {

	Email 	string `json:"email" validate:"required,email"`
}
