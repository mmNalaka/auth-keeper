package models

import "time"

type Role struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Disabled bool   `json:"disabled" gorm:"default:false"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
