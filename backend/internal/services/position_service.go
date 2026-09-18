package services

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/repositories"
	"github.com/google/uuid"
)

type PositionService struct {
	repo *repositories.PositionRepository
}

func NewPositionService(r *repositories.PositionRepository) *PositionService {
	return &PositionService{r}
}
func (s *PositionService) Create(x *models.Position) error            { return s.repo.Create(x) }
func (s *PositionService) List() ([]models.Position, error)           { return s.repo.List() }
func (s *PositionService) Get(id uuid.UUID) (*models.Position, error) { return s.repo.Get(id) }
func (s *PositionService) Update(x *models.Position) error            { return s.repo.Update(x) }
func (s *PositionService) Delete(id uuid.UUID) error                  { return s.repo.Delete(id) }
func (s *PositionService) ListByDepartment(departmentID uuid.UUID) ([]models.Position, error) {
	return s.repo.ListByDepartment(departmentID)
}
