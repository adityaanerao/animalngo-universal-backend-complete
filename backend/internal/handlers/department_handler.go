package handlers

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/services"
	"animalngo-universal-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DepartmentHandler struct{ s *services.DepartmentService }

func NewDepartmentHandler(s *services.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{s}
}
type CreateDepartmentRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// @Summary Create department
// @Description Create a new department
// @Tags departments
// @Accept json
// @Produce json
// @Param request body CreateDepartmentRequest true "Department request"
// @Success 201 {object} models.Department
// @Failure 400 {object} response.ErrorResponse
// @Router /departments [post]
// @Security BearerAuth
func (h *DepartmentHandler) Create(c *gin.Context) {
	var req CreateDepartmentRequest
	if e := c.ShouldBindJSON(&req); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	x := models.Department{
		Name:        req.Name,
		Description: req.Description,
	}
	if e := h.s.Create(&x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 201, "department created", x)
}
// @Summary List departments
// @Description Get a list of departments
// @Tags departments
// @Produce json
// @Success 200 {array} models.Department
// @Failure 500 {object} response.ErrorResponse
// @Router /departments [get]
// @Security BearerAuth
func (h *DepartmentHandler) List(c *gin.Context) {
	x, e := h.s.List()
	if e != nil {
		response.Error(c, 500, e.Error())
		return
	}
	response.Success(c, 200, "departments", x)
}
// @Summary Get department
// @Description Get department by ID
// @Tags departments
// @Produce json
// @Param id path string true "Department ID"
// @Success 200 {object} models.Department
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /departments/{id} [get]
// @Security BearerAuth
func (h *DepartmentHandler) Get(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	x, e := h.s.Get(id)
	if e != nil {
		response.Error(c, 404, "department not found")
		return
	}
	response.Success(c, 200, "department", x)
}
// @Summary Update department
// @Description Update department by ID
// @Tags departments
// @Accept json
// @Produce json
// @Param id path string true "Department ID"
// @Param request body models.Department true "Department request"
// @Success 200 {object} models.Department
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /departments/{id} [put]
// @Security BearerAuth
func (h *DepartmentHandler) Update(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	x, e := h.s.Get(id)
	if e != nil {
		response.Error(c, 404, "department not found")
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
	response.Success(c, 200, "department updated", x)
}
// @Summary Delete department
// @Description Delete department by ID
// @Tags departments
// @Param id path string true "Department ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /departments/{id} [delete]
// @Security BearerAuth
func (h *DepartmentHandler) Delete(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	if e = h.s.Delete(id); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 200, "department deleted", nil)
}

