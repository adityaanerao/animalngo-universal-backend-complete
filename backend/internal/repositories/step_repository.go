package repositories

import (
	"animalngo-universal-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StepRepository struct{ db *gorm.DB }

func NewStepRepository(db *gorm.DB) *StepRepository { return &StepRepository{db} }

func (r *StepRepository) Create(x *models.WorkflowStep) error {
	return r.db.Create(x).Error
}

func (r *StepRepository) ListByWorkflow(workflowID uuid.UUID) ([]models.WorkflowStep, error) {
	var x []models.WorkflowStep
	e := r.db.Where("workflow_id = ?", workflowID).Order("step_order ASC").Find(&x).Error
	return x, e
}

func (r *StepRepository) Get(id uuid.UUID) (*models.WorkflowStep, error) {
	var x models.WorkflowStep
	e := r.db.First(&x, "id = ?", id).Error
	return &x, e
}

func (r *StepRepository) Update(x *models.WorkflowStep) error {
	return r.db.Save(x).Error
}

func (r *StepRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.WorkflowStep{}, "id = ?", id).Error
}
