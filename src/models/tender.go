package models

import (
    "time"

    "github.com/google/uuid"
)

type Tender struct {
    ID             uuid.UUID   `gorm:"type:uuid;primaryKey"`
    Name           string      `gorm:"not null"`
    Description    string
    ServiceType    string      `gorm:"not null"`
    Status         TenderStatus `gorm:"not null"`
    OrganizationID uuid.UUID   `gorm:"type:uuid;not null"`
    CreatorUsername string      `gorm:"not null"`
    CreatedAt      time.Time `gorm:"autoCreateTime"`
    UpdatedAt      time.Time `gorm:"autoUpdateTime"`
    Version        int         `gorm:"not null"`
}
