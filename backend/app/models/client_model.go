package models

import "github.com/google/uuid"

type Client struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Secret    string    `json:"secret"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}
