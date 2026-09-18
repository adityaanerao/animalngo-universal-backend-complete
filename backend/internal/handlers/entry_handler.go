package handlers

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/services"
	"animalngo-universal-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EntryHandler struct{ s *services.EntryService }

func NewEntryHandler(s *services.EntryService) *EntryHandler { return &EntryHandler{s} }
// @Summary Submit entry
// @Description Submit data for a specific workflow step
// @Tags entries
// @Accept json
// @Produce json
// @Param id path string true "Workflow ID"
// @Param step_id path string true "Step ID"
// @Param request body models.DynamicEntry true "Entry request"
// @Success 201 {object} models.DynamicEntry
// @Failure 400 {object} response.ErrorResponse
// @Router /workflows/{id}/steps/{step_id}/submit [post]
// @Security BearerAuth
func (h *EntryHandler) Create(c *gin.Context) {
	workflowID, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid workflow id")
		return
	}
	stepID, e := uuid.Parse(c.Param("step_id"))
	if e != nil {
		response.Error(c, 400, "invalid step id")
		return
	}

	var x models.DynamicEntry
	if e := c.ShouldBindJSON(&x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	x.WorkflowID = workflowID
	x.StepID = stepID
	x.SubmittedBy = uuid.MustParse(c.GetString("user_id"))
	if e := h.s.Create(&x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 201, "entry submitted", x)
}
// @Summary List entries
// @Description Get a list of entries
// @Tags entries
// @Produce json
// @Success 200 {array} models.DynamicEntry
// @Failure 500 {object} response.ErrorResponse
// @Router /entries [get]
// @Security BearerAuth
func (h *EntryHandler) List(c *gin.Context) {
	x, e := h.s.List()
	if e != nil {
		response.Error(c, 500, e.Error())
		return
	}
	response.Success(c, 200, "entries", x)
}
// @Summary Get entry
// @Description Get entry by ID
// @Tags entries
// @Produce json
// @Param id path string true "Entry ID"
// @Success 200 {object} models.DynamicEntry
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /entries/{id} [get]
// @Security BearerAuth
func (h *EntryHandler) Get(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	x, e := h.s.Get(id)
	if e != nil {
		response.Error(c, 404, "entry not found")
		return
	}
	response.Success(c, 200, "entry", x)
}

