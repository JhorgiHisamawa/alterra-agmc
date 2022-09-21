package models

import (
	"time"
)

type User struct {
	Id        int       `json:"id" form:"id" gorm:"primaryKey"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email" validate:"required"`
	Password  string    `json:"password" form:"password" validate:"required"`
	Role      string    `json:"role" form:"role"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}

type UserResponse struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}
