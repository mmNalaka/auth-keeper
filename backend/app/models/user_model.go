package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           string `json:"id" gorm:"type:uuid;primary_key"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	IsSuperAdmin bool   `json:"-" gorm:"default:false"`
	ProfileImage string `json:"profile_image"`

	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`
	InvitedAt   *time.Time `json:"invited_at,omitempty"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	ConfirmationToken  string     `json:"-"`
	ConfirmationSentAt *time.Time `json:"confirmation_sent_at,omitempty"`

	ResetPasswordToken  string     `json:"-"`
	ResetPasswordSentAt *time.Time `json:"reset_password_sent_at,omitempty"`

	Blocked   bool       `json:"blocked" gorm:"default:false"`
	BlockedAt *time.Time `json:"blocked_at,omitempty"`

	LastSignInAt *time.Time `json:"last_sign_in_at,omitempty"`
}
