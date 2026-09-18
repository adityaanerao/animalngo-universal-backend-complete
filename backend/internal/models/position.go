package models

import (
	"time"

	"github.com/google/uuid"
)

type Position struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	DepartmentID uuid.UUID `gorm:"type:uuid;not null;index" json:"department_id"`

	Name        string `gorm:"type:varchar(150);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Department *Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Users      []User      `gorm:"foreignKey:PositionID" json:"users,omitempty"`
}
