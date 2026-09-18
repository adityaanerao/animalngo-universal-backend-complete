package services

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/repositories"
	"github.com/google/uuid"
)

type DepartmentService struct {
	repo *repositories.DepartmentRepository
}

func NewDepartmentService(r *repositories.DepartmentRepository) *DepartmentService {
	return &DepartmentService{r}
}
func (s *DepartmentService) Create(x *models.Department) error            { return s.repo.Create(x) }
func (s *DepartmentService) List() ([]models.Department, error)           { return s.repo.List() }
func (s *DepartmentService) Get(id uuid.UUID) (*models.Department, error) { return s.repo.Get(id) }
func (s *DepartmentService) Update(x *models.Department) error            { return s.repo.Update(x) }
func (s *DepartmentService) Delete(id uuid.UUID) error                    { return s.repo.Delete(id) }
