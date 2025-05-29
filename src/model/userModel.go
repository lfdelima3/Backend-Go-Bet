package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100) json:"name" validate:"required"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(100);not null" json:"-" validate:"required,min=6"`
	//Data_nasc
	//Time_Torcedor
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
