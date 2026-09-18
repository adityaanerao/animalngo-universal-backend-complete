package repositories

import (
	"animalngo-universal-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntryRepository struct{ db *gorm.DB }

func NewEntryRepository(db *gorm.DB) *EntryRepository          { return &EntryRepository{db} }
func (r *EntryRepository) Create(x *models.DynamicEntry) error { return r.db.Create(x).Error }
func (r *EntryRepository) List() ([]models.DynamicEntry, error) {
	var x []models.DynamicEntry
	e := r.db.Order("created_at desc").Find(&x).Error
	return x, e
}
func (r *EntryRepository) Get(id uuid.UUID) (*models.DynamicEntry, error) {
	var x models.DynamicEntry
	e := r.db.First(&x, "id = ?", id).Error
	return &x, e
}
