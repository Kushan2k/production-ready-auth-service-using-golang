package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;uniqueIndex;not null"`
	Password string `gorm:"size:8,not null"`
}


type SwaggerUser struct {
    ID        uint      `json:"id" gorm:"primarykey"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at,omitempty"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
}
