package services

import (
	"animalngo-universal-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RuleService struct{ db *gorm.DB }

func NewRuleService(db *gorm.DB) *RuleService          { return &RuleService{db} }
func (s *RuleService) Create(x *models.StepRule) error { return s.db.Create(x).Error }
func (s *RuleService) List(stepID uuid.UUID) ([]models.StepRule, error) {
	var x []models.StepRule
	e := s.db.Where("step_id = ?", stepID).Find(&x).Error
	return x, e
}
func (s *RuleService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.StepRule{}, "id = ?", id).Error
}
