package repositories

import (
	"animalngo-universal-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkflowRepository struct{ db *gorm.DB }

func NewWorkflowRepository(db *gorm.DB) *WorkflowRepository   { return &WorkflowRepository{db} }
func (r *WorkflowRepository) Create(x *models.Workflow) error { return r.db.Create(x).Error }
func (r *WorkflowRepository) List() ([]models.Workflow, error) {
	var x []models.Workflow
	e := r.db.Preload("Steps.Fields").Preload("Steps.Rules").Find(&x).Error
	return x, e
}
func (r *WorkflowRepository) Get(id uuid.UUID) (*models.Workflow, error) {
	var x models.Workflow
	e := r.db.Preload("Steps.Fields").Preload("Steps.Rules").First(&x, "id = ?", id).Error
	return &x, e
}
func (r *WorkflowRepository) Update(x *models.Workflow) error { return r.db.Save(x).Error }
func (r *WorkflowRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Workflow{}, "id = ?", id).Error
}
