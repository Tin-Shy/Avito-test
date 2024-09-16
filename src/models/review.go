package models

import "github.com/google/uuid"

type Review struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"createdAt"`
}
