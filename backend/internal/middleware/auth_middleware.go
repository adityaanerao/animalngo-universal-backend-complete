package middleware

import (
	"animalngo-universal-backend/config"
	jwtutil "animalngo-universal-backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		p := strings.Fields(h)
		if len(p) != 2 || strings.ToLower(p[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Bearer token required"})
			return
		}
		cl, e := jwtutil.ParseToken(p[1], cfg.JWTSecret)
		if e != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid or expired token"})
			return
		}
		c.Set("user_id", cl.UserID)
		c.Set("position_id", cl.PositionID)
		c.Next()
	}
}
