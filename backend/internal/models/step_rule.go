package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type StepRule struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StepID uuid.UUID `gorm:"type:uuid;not null;index" json:"step_id"`

	ConditionIf datatypes.JSON `gorm:"type:jsonb" json:"condition_if" swaggertype:"object"`
	ActionThen  datatypes.JSON `gorm:"type:jsonb" json:"action_then" swaggertype:"object"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
