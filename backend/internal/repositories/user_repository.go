package repositories

import (
	"animalngo-universal-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) *UserRepository   { return &UserRepository{db} }
func (r *UserRepository) Create(x *models.User) error { return r.db.Create(x).Error }
func (r *UserRepository) ByMobileNumber(mobile string) (*models.User, error) {
	var x models.User
	e := r.db.Preload("Position").Where("mobile_number = ?", mobile).First(&x).Error
	return &x, e
}
func (r *UserRepository) ByID(id uuid.UUID) (*models.User, error) {
	var x models.User
	e := r.db.Preload("Position").First(&x, "id = ?", id).Error
	return &x, e
}
func (r *UserRepository) Update(x *models.User) error {
	return r.db.Save(x).Error
}
