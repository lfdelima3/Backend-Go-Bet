package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"-" validate:"required,min=8,strong_password"`
	Role      string    `json:"role" validate:"required,oneof=user admin"`
	Balance   float64   `json:"balance" validate:"required,min=0"`
	Status    string    `json:"status" validate:"required,oneof=active inactive blocked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreate struct {
	Name     string  `json:"name" validate:"required,min=3,max=100"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8,strong_password"`
	Role     string  `json:"role" validate:"required,oneof=user admin"`
	Balance  float64 `json:"balance" validate:"required,min=0"`
}

type UserUpdate struct {
	Name     string  `json:"name" validate:"omitempty,min=3,max=100"`
	Email    string  `json:"email" validate:"omitempty,email"`
	Password string  `json:"password" validate:"omitempty,min=8,strong_password"`
	Role     string  `json:"role" validate:"omitempty,oneof=user admin"`
	Balance  float64 `json:"balance" validate:"omitempty,min=0"`
	Status   string  `json:"status" validate:"omitempty,oneof=active inactive blocked"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Balance   float64   `json:"balance"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
