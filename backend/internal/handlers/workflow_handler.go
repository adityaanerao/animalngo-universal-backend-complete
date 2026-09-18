package handlers

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/services"
	"animalngo-universal-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkflowHandler struct{ s *services.WorkflowService }

func NewWorkflowHandler(s *services.WorkflowService) *WorkflowHandler { return &WorkflowHandler{s} }
type CreateWorkflowRequest struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	DepartmentID string `json:"department_id" binding:"required"`
}

// @Summary Create workflow
// @Description Create a new workflow
// @Tags workflows
// @Accept json
// @Produce json
// @Param request body CreateWorkflowRequest true "Workflow request"
// @Success 201 {object} models.Workflow
// @Failure 400 {object} response.ErrorResponse
// @Router /workflows [post]
// @Security BearerAuth
func (h *WorkflowHandler) Create(c *gin.Context) {
	var req CreateWorkflowRequest
	if e := c.ShouldBindJSON(&req); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	depID, e := uuid.Parse(req.DepartmentID)
	if e != nil {
		response.Error(c, 400, "invalid department id")
		return
	}
	x := models.Workflow{
		Name:         req.Name,
		Description:  req.Description,
		DepartmentID: depID,
		CreatedBy:    uuid.MustParse(c.GetString("user_id")),
	}
	if e := h.s.Create(&x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 201, "workflow created", x)
}
// @Summary List workflows
// @Description Get a list of workflows
// @Tags workflows
// @Produce json
// @Success 200 {array} models.Workflow
// @Failure 500 {object} response.ErrorResponse
// @Router /workflows [get]
// @Security BearerAuth
func (h *WorkflowHandler) List(c *gin.Context) {
	x, e := h.s.List()
	if e != nil {
		response.Error(c, 500, e.Error())
		return
	}
	response.Success(c, 200, "workflows", x)
}
// @Summary Get workflow
// @Description Get workflow by ID
// @Tags workflows
// @Produce json
// @Param id path string true "Workflow ID"
// @Success 200 {object} models.Workflow
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /workflows/{id} [get]
// @Security BearerAuth
func (h *WorkflowHandler) Get(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	x, e := h.s.Get(id)
	if e != nil {
		response.Error(c, 404, "workflow not found")
		return
	}
	response.Success(c, 200, "workflow", x)
}
// @Summary Update workflow
// @Description Update workflow by ID
// @Tags workflows
// @Accept json
// @Produce json
// @Param id path string true "Workflow ID"
// @Param request body models.Workflow true "Workflow request"
// @Success 200 {object} models.Workflow
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /workflows/{id} [put]
// @Security BearerAuth
func (h *WorkflowHandler) Update(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	x, e := h.s.Get(id)
	if e != nil {
		response.Error(c, 404, "workflow not found")
		return
	}
	if e = c.ShouldBindJSON(x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	x.ID = id
	if e = h.s.Update(x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 200, "workflow updated", x)
}
// @Summary Delete workflow
// @Description Delete workflow by ID
// @Tags workflows
// @Param id path string true "Workflow ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /workflows/{id} [delete]
// @Security BearerAuth
func (h *WorkflowHandler) Delete(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	if e = h.s.Delete(id); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 200, "workflow deleted", nil)
}

