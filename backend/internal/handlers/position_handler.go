package handlers

import (
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/services"
	"animalngo-universal-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PositionHandler struct{ s *services.PositionService }

func NewPositionHandler(s *services.PositionService) *PositionHandler { return &PositionHandler{s} }
type CreatePositionRequest struct {
	DepartmentID string `json:"department_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
}

// @Summary Create position
// @Description Create a new position
// @Tags positions
// @Accept json
// @Produce json
// @Param request body CreatePositionRequest true "Position request"
// @Success 201 {object} models.Position
// @Failure 400 {object} response.ErrorResponse
// @Router /positions [post]
// @Security BearerAuth
func (h *PositionHandler) Create(c *gin.Context) {
	var req CreatePositionRequest
	if e := c.ShouldBindJSON(&req); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	depID, e := uuid.Parse(req.DepartmentID)
	if e != nil {
		response.Error(c, 400, "invalid department id")
		return
	}
	x := models.Position{
		DepartmentID: depID,
		Name:         req.Name,
		Description:  req.Description,
	}
	if e := h.s.Create(&x); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 201, "position created", x)
}
// @Summary List positions
// @Description Get a list of positions
// @Tags positions
// @Produce json
// @Success 200 {array} models.Position
// @Failure 500 {object} response.ErrorResponse
// @Router /positions [get]
// @Security BearerAuth
func (h *PositionHandler) List(c *gin.Context) {
	x, e := h.s.List()
	if e != nil {
		response.Error(c, 500, e.Error())
		return
	}
	response.Success(c, 200, "positions", x)
}

// @Summary List positions by department
// @Description Get a list of positions by department ID
// @Tags departments
// @Produce json
// @Param id path string true "Department ID"
// @Success 200 {array} models.Position
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /departments/{id}/positions [get]
// @Security BearerAuth
func (h *PositionHandler) ListByDepartment(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid department id")
		return
	}
	x, e := h.s.ListByDepartment(id)
	if e != nil {
		response.Error(c, 500, e.Error())
		return
	}
	response.Success(c, 200, "positions", x)
}
// @Summary Get position
// @Description Get position by ID
// @Tags positions
// @Produce json
// @Param id path string true "Position ID"
// @Success 200 {object} models.Position
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /positions/{id} [get]
// @Security BearerAuth
func (h *PositionHandler) Get(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	x, e := h.s.Get(id)
	if e != nil {
		response.Error(c, 404, "position not found")
		return
	}
	response.Success(c, 200, "position", x)
}
// @Summary Update position
// @Description Update position by ID
// @Tags positions
// @Accept json
// @Produce json
// @Param id path string true "Position ID"
// @Param request body models.Position true "Position request"
// @Success 200 {object} models.Position
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /positions/{id} [put]
// @Security BearerAuth
func (h *PositionHandler) Update(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	x, e := h.s.Get(id)
	if e != nil {
		response.Error(c, 404, "position not found")
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
	response.Success(c, 200, "position updated", x)
}
// @Summary Delete position
// @Description Delete position by ID
// @Tags positions
// @Param id path string true "Position ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /positions/{id} [delete]
// @Security BearerAuth
func (h *PositionHandler) Delete(c *gin.Context) {
	id, e := uuid.Parse(c.Param("id"))
	if e != nil {
		response.Error(c, 400, "invalid id")
		return
	}
	if e = h.s.Delete(id); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, 200, "position deleted", nil)
}

