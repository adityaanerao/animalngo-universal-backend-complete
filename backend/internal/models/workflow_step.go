package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkflowStep struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	WorkflowID uuid.UUID `gorm:"type:uuid;not null;index" json:"workflow_id"`
	PositionID uuid.UUID `gorm:"type:uuid;not null;index" json:"position_id"`

	Name        string `gorm:"type:varchar(200);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	StepOrder int  `gorm:"not null" json:"step_order"`
	IsActive  bool `gorm:"default:true" json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Workflow *Workflow `gorm:"foreignKey:WorkflowID" json:"workflow,omitempty"`
	Position *Position `gorm:"foreignKey:PositionID" json:"position,omitempty"`

	Fields []FormField `gorm:"foreignKey:StepID;constraint:OnDelete:CASCADE" json:"fields,omitempty"`
	Rules  []StepRule  `gorm:"foreignKey:StepID;constraint:OnDelete:CASCADE" json:"rules,omitempty"`
}
