package handlers

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/services"
	"animalngo-universal-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RuleHandler struct{ s *services.RuleService }

func NewRuleHandler(s *services.RuleService) *RuleHandler { return &RuleHandler{s} }

// @Summary Create rule
// @Description Create a new rule for a step
// @Tags step-rules
// @Accept json
// @Produce json
// @Param id path string true "Step ID"
// @Param request body models.StepRule true "Rule request"
// @Success 201 {object} models.StepRule
// @Failure 400 {object} response.ErrorResponse
// @Router /steps/{id}/rules [post]
// @Security BearerAuth
func (h *RuleHandler) Create(c *gin.Context) {
	stepID, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid step id")
		return
	}
	var x models.StepRule
	if e := c.ShouldBindJSON(&x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	x.StepID = stepID
	if e := h.s.Create(&x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 201, "rule created", x)
}
