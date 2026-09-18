package services

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/repositories"
	"github.com/google/uuid"
)

type StepService struct {
	repo *repositories.StepRepository
}

func NewStepService(r *repositories.StepRepository) *StepService {
	return &StepService{r}
}

func (s *StepService) Create(x *models.WorkflowStep) error {
	return s.repo.Create(x)
}

func (s *StepService) ListByWorkflow(workflowID uuid.UUID) ([]models.WorkflowStep, error) {
	return s.repo.ListByWorkflow(workflowID)
}

func (s *StepService) Get(id uuid.UUID) (*models.WorkflowStep, error) {
	return s.repo.Get(id)
}

func (s *StepService) Update(x *models.WorkflowStep) error {
	return s.repo.Update(x)
}

func (s *StepService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
