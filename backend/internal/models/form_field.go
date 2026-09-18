package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type FormField struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StepID uuid.UUID `gorm:"type:uuid;not null;index" json:"step_id"`

	FieldKey    string `gorm:"type:varchar(100);not null" json:"field_key"`
	Label       string `gorm:"type:varchar(200);not null" json:"label"`
	FieldType   string `gorm:"type:varchar(50);not null" json:"field_type"`
	Placeholder string `gorm:"type:varchar(255)" json:"placeholder"`

	IsRequired bool `gorm:"default:false" json:"is_required"`
	FieldOrder int  `gorm:"not null" json:"field_order"`

	Options    datatypes.JSON `gorm:"type:jsonb" json:"options,omitempty" swaggertype:"object"`
	Validation datatypes.JSON `gorm:"type:jsonb" json:"validation,omitempty" swaggertype:"object"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
