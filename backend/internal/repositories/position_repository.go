package repositories

import (
	"animalngo-universal-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PositionRepository struct{ db *gorm.DB }

func NewPositionRepository(db *gorm.DB) *PositionRepository   { return &PositionRepository{db} }
func (r *PositionRepository) Create(x *models.Position) error { return r.db.Create(x).Error }
func (r *PositionRepository) List() ([]models.Position, error) {
	var x []models.Position
	e := r.db.Preload("Department").Find(&x).Error
	return x, e
}
func (r *PositionRepository) Get(id uuid.UUID) (*models.Position, error) {
	var x models.Position
	e := r.db.Preload("Department").First(&x, "id = ?", id).Error
	return &x, e
}
func (r *PositionRepository) Update(x *models.Position) error { return r.db.Save(x).Error }
func (r *PositionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Position{}, "id = ?", id).Error
}
func (r *PositionRepository) ListByDepartment(departmentID uuid.UUID) ([]models.Position, error) {
	var x []models.Position
	e := r.db.Where("department_id = ?", departmentID).Find(&x).Error
	return x, e
}
