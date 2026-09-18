package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type DynamicEntry struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	WorkflowID  uuid.UUID `gorm:"type:uuid;not null;index" json:"workflow_id"`
	StepID      uuid.UUID `gorm:"type:uuid;not null;index" json:"step_id"`
	SubmittedBy uuid.UUID `gorm:"type:uuid;not null;index" json:"submitted_by"`

	Data datatypes.JSON `gorm:"type:jsonb;not null" json:"data" swaggertype:"object"`

	Status string `gorm:"type:varchar(30);default:'submitted'" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
