package models

import (
    "time"

    "github.com/google/uuid"
)

type Bid struct {
    ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
    Amount         float64
    UserID         uuid.UUID
    TenderID       uuid.UUID
    CreatorUsername string
    CreatedAt      time.Time `gorm:"autoCreateTime"`
    UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
