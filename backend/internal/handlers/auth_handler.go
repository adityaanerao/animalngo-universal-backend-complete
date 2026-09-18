package handlers

import (
	"animalngo-universal-backend/internal/services"
	"animalngo-universal-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type AuthHandler struct{ s *services.AuthService }

func NewAuthHandler(s *services.AuthService) *AuthHandler { return &AuthHandler{s} }

type sendSignupOTPRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required"`
	FullName     string `json:"full_name" binding:"required"`
}

type sendLoginOTPRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required"`
}

type verifyOTPRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required"`
	OTP          string `json:"otp" binding:"required"`
}

// @Summary Send Signup OTP
// @Description Send an OTP for user registration
// @Tags auth
// @Accept json
// @Produce json
// @Param request body sendSignupOTPRequest true "Send OTP Request"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/send-otp/signup [post]
func (h *AuthHandler) SendSignupOTP(c *gin.Context) {
	var r sendSignupOTPRequest
	if e := c.ShouldBindJSON(&r); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	if e := h.s.SendSignupOTP(r.MobileNumber, r.FullName); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, http.StatusOK, "OTP sent successfully", nil)
}

// @Summary Send Login OTP
// @Description Send an OTP for user login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body sendLoginOTPRequest true "Send Login OTP Request"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/send-otp/login [post]
func (h *AuthHandler) SendLoginOTP(c *gin.Context) {
	var r sendLoginOTPRequest
	if e := c.ShouldBindJSON(&r); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	if e := h.s.SendLoginOTP(r.MobileNumber); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	response.Success(c, http.StatusOK, "OTP sent successfully", nil)
}

// @Summary Verify OTP
// @Description Verify OTP and login/register
// @Tags auth
// @Accept json
// @Produce json
// @Param request body verifyOTPRequest true "Verify OTP Request"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /auth/verify-otp [post]
func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var r verifyOTPRequest
	if e := c.ShouldBindJSON(&r); e != nil {
		response.Error(c, 400, e.Error())
		return
	}
	u, t, e := h.s.VerifyOTP(r.MobileNumber, r.OTP)
	if e != nil {
		response.Error(c, 401, e.Error())
		return
	}
	response.Success(c, http.StatusOK, "Verification successful", gin.H{"user": u, "token": t})
}
// @Summary Get current user
// @Description Get current authenticated user details
// @Tags auth
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /auth/me [get]
// @Security BearerAuth
func (h *AuthHandler) Me(c *gin.Context) {
	id, e := uuid.Parse(c.GetString("user_id"))
	if e != nil {
		response.Error(c, 401, "invalid user context")
		return
	}
	u, e := h.s.ByID(id)
	if e != nil {
		response.Error(c, 404, "user not found")
		return
	}
	response.Success(c, 200, "current user", u)
}
