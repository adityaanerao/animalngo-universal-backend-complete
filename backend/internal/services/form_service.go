package services

import (
	"animalngo-universal-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FormService struct{ db *gorm.DB }

func NewFormService(db *gorm.DB) *FormService           { return &FormService{db} }
func (s *FormService) Create(x *models.FormField) error { return s.db.Create(x).Error }
func (s *FormService) List(stepID uuid.UUID) ([]models.FormField, error) {
	var x []models.FormField
	e := s.db.Where("step_id = ?", stepID).Order("field_order").Find(&x).Error
	return x, e
}
func (s *FormService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.FormField{}, "id = ?", id).Error
}
