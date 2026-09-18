package repositories

import (
	"animalngo-universal-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentRepository struct{ db *gorm.DB }

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository   { return &DepartmentRepository{db} }
func (r *DepartmentRepository) Create(x *models.Department) error { return r.db.Create(x).Error }
func (r *DepartmentRepository) List() ([]models.Department, error) {
	var x []models.Department
	e := r.db.Preload("Positions").Find(&x).Error
	return x, e
}
func (r *DepartmentRepository) Get(id uuid.UUID) (*models.Department, error) {
	var x models.Department
	e := r.db.Preload("Positions").First(&x, "id = ?", id).Error
	return &x, e
}
func (r *DepartmentRepository) Update(x *models.Department) error { return r.db.Save(x).Error }
func (r *DepartmentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Department{}, "id = ?", id).Error
}
