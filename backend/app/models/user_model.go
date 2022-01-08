package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Email        string    `json:"email" gorm:"unique"`
	Password     string    `json:"password"`
	Role         string    `json:"role" gorm:"default:'guest'"`
	IsSuperAdmin bool      `json:"-" gorm:"default:false"`
	ProfileImage string    `json:"profile_image"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	DateOfBirth  time.Time `json:"date_of_birth,omitempty"`

	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`
	InvitedAt   *time.Time `json:"invited_at,omitempty"`

	ConfirmationToken  string     `json:"-"`
	ConfirmationSentAt *time.Time `json:"confirmation_sent_at,omitempty"`

	ResetPasswordToken  string     `json:"-"`
	ResetPasswordSentAt *time.Time `json:"reset_password_sent_at,omitempty"`

	Blocked   bool       `json:"blocked" gorm:"default:false"`
	BlockedAt *time.Time `json:"blocked_at,omitempty"`

	LastSignInAt *time.Time `json:"last_sign_in_at,omitempty"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
