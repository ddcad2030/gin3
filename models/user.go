package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	// ID       uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	// ID       uuid.UUID `gorm:"primary_key; unique; type:uuid; default:uuid_generate_v4()"`

	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(255); not null"`
	Email    string `gorm:"uniqueIndex; not null"`
	Password string `gorm:"type:varchar(255); not null"`
	// CreatedAt time.Time `gorm:"not null"`
	// UpdatedAt time.Time `gorm:"not null"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreate struct {
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUpdate struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
