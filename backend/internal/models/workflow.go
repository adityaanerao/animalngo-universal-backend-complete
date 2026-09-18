package models

import (
	"time"

	"github.com/google/uuid"
)

type Workflow struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	Name         string `gorm:"type:varchar(200);not null" json:"name"`
	Description  string `gorm:"type:text" json:"description"`
	DepartmentID uuid.UUID `gorm:"type:uuid;not null;index" json:"department_id"`

	TargetTableType string `gorm:"type:varchar(50);default:'DYNAMIC_ENTRY'" json:"target_table_type"`

	Status string `gorm:"type:varchar(30);default:'draft';index" json:"status"`

	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Steps []WorkflowStep `gorm:"foreignKey:WorkflowID;constraint:OnDelete:CASCADE" json:"steps,omitempty"`
}
