package services

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/repositories"
	"github.com/google/uuid"
)

type EntryService struct{ repo *repositories.EntryRepository }

func NewEntryService(r *repositories.EntryRepository) *EntryService    { return &EntryService{r} }
func (s *EntryService) Create(x *models.DynamicEntry) error            { return s.repo.Create(x) }
func (s *EntryService) List() ([]models.DynamicEntry, error)           { return s.repo.List() }
func (s *EntryService) Get(id uuid.UUID) (*models.DynamicEntry, error) { return s.repo.Get(id) }
