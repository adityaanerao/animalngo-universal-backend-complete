package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	PositionID *uuid.UUID `gorm:"type:uuid;index" json:"position_id"`

	FullName     string `gorm:"type:varchar(150);not null" json:"full_name"`
	MobileNumber string `gorm:"type:varchar(20);not null;uniqueIndex" json:"mobile_number"`
	SessionToken string `gorm:"type:varchar(255)" json:"session_token"`

	IsActive bool `gorm:"default:true" json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Position *Position `gorm:"foreignKey:PositionID" json:"position,omitempty"`
}
