package models

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	WorkflowID uuid.UUID `gorm:"type:uuid;not null;index" json:"workflow_id"`

	FirstName  string    `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName   string    `gorm:"type:varchar(100);not null" json:"last_name"`
	DOB        time.Time `json:"dob"`
	BloodGroup string    `gorm:"type:varchar(10)" json:"blood_group"`
	Phone      string    `gorm:"type:varchar(20)" json:"phone"`
	Email      string    `gorm:"type:varchar(200)" json:"email"`
	Address    string    `gorm:"type:text" json:"address"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
