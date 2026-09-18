package services

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/repositories"
	"github.com/google/uuid"
)

type WorkflowService struct {
	repo *repositories.WorkflowRepository
}

func NewWorkflowService(r *repositories.WorkflowRepository) *WorkflowService {
	return &WorkflowService{r}
}
func (s *WorkflowService) Create(x *models.Workflow) error            { return s.repo.Create(x) }
func (s *WorkflowService) List() ([]models.Workflow, error)           { return s.repo.List() }
func (s *WorkflowService) Get(id uuid.UUID) (*models.Workflow, error) { return s.repo.Get(id) }
func (s *WorkflowService) Update(x *models.Workflow) error            { return s.repo.Update(x) }
func (s *WorkflowService) Delete(id uuid.UUID) error                  { return s.repo.Delete(id) }
