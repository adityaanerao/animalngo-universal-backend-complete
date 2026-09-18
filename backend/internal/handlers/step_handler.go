package handlers

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/services"
	"animalngo-universal-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StepHandler struct{ s *services.StepService }

func NewStepHandler(s *services.StepService) *StepHandler { return &StepHandler{s} }

type CreateWorkflowStepRequest struct {
	Name             string `json:"name" binding:"required"`
	PositionEntityID string `json:"position_entity_id" binding:"required"`
	StepOrder        int    `json:"step_order" binding:"required"`
}

// @Summary Create step
// @Description Create a new workflow step
// @Tags workflow-steps
// @Accept json
// @Produce json
// @Param id path string true "Workflow ID"
// @Param request body CreateWorkflowStepRequest true "Step request"
// @Success 201 {object} models.WorkflowStep
// @Failure 400 {object} response.ErrorResponse
// @Router /workflows/{id}/steps [post]
// @Security BearerAuth
func (h *StepHandler) Create(c *gin.Context) {
	workflowID, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid workflow id")
		return
	}
	var req CreateWorkflowStepRequest
	if e := c.ShouldBindJSON(&req); e != nil {
		response.Error(c, 400, e.Error())
		return
	}

	positionID, e := uuid.Parse(req.PositionEntityID)
	if e != nil {
		response.Error(c, 400, "invalid position_entity_id")
		return
	}

	x := models.WorkflowStep{
		WorkflowID: workflowID,
		Name:       req.Name,
		PositionID: positionID,
		StepOrder:  req.StepOrder,
	}

	if e := h.s.Create(&x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 201, "step created", x)
}

type ReorderRequest struct {
	Steps []struct {
		ID        uuid.UUID `json:"id" binding:"required"`
		StepOrder int       `json:"step_order" binding:"required"`
	} `json:"steps" binding:"required"`
}

// @Summary Reorder steps
// @Description Reorder workflow steps
// @Tags workflow-steps
// @Accept json
// @Produce json
// @Param id path string true "Workflow ID"
// @Param request body ReorderRequest true "Reorder request"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /workflows/{id}/steps/reorder [put]
// @Security BearerAuth
func (h *StepHandler) Reorder(c *gin.Context) {
	_, e := uuid.Parse(c.Param("id")) // validate it's uuid
	if e != nil {
		response.Error(c, 400, "invalid workflow id")
		return
	}
	var req ReorderRequest
	if e := c.ShouldBindJSON(&req); e != nil {
		response.Error(c, 400, e.Error())
		return
	}

	for _, reqStep := range req.Steps {
		step, e := h.s.Get(reqStep.ID)
		if e == nil {
			step.StepOrder = reqStep.StepOrder
			_ = h.s.Update(step)
		}
	}
	response.Success(c, 200, "steps reordered", nil)
}
