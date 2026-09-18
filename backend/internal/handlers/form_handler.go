package handlers

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/services"
	"animalngo-universal-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FormHandler struct{ s *services.FormService }

func NewFormHandler(s *services.FormService) *FormHandler { return &FormHandler{s} }
// @Summary Create form field
// @Description Create a new form field for a step
// @Tags form-fields
// @Accept json
// @Produce json
// @Param id path string true "Step ID"
// @Param request body models.FormField true "Form request"
// @Success 201 {object} models.FormField
// @Failure 400 {object} response.ErrorResponse
// @Router /steps/{id}/fields [post]
// @Security BearerAuth
func (h *FormHandler) Create(c *gin.Context) {
	stepID, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid step id")
		return
	}
	var x models.FormField
	if e := c.ShouldBindJSON(&x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	x.StepID = stepID
	if e := h.s.Create(&x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 201, "form field created", x)
}
// @Summary List form-fields
// @Description Get a list of form-fields
// @Tags form-fields
// @Produce json
// @Param stepID path string true "Step ID"
// @Success 200 {array} models.FormField
// @Failure 500 {object} response.ErrorResponse
// @Router /form-fields/step/{stepID} [get]
// @Security BearerAuth
func (h *FormHandler) List(c *gin.Context) {
	id, e := uuid.Parse(c.Param("stepID"))
	if e != nil {
		response.Error(c, 400, "invalid step id")
		return
	}
	x, e := h.s.List(id)
	if e != nil {
		response.Error(c, 500, e.Error())
		return
	}
	response.Success(c, 200, "form fields", x)
}
// @Summary Delete form
// @Description Delete form by ID
// @Tags form-fields
// @Param id path string true "Form ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /form-fields/{id} [delete]
// @Security BearerAuth
func (h *FormHandler) Delete(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	if e = h.s.Delete(id); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 200, "form field deleted", nil)
}

