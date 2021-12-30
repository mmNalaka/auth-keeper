package models

import (
	"time"
)

type RefreshToken struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Token   string `json:"token"`
	UserID  string `json:"user_id"`
	Revoked bool   `json:"revoked"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
