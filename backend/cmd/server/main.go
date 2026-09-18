package main

import (
	"animalngo-universal-backend/config"
	_ "animalngo-universal-backend/docs"
	"animalngo-universal-backend/internal/handlers"
	"animalngo-universal-backend/internal/middleware"
	"animalngo-universal-backend/internal/repositories"
	"animalngo-universal-backend/internal/services"
	"animalngo-universal-backend/pkg/database"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Animal NGO Universal Backend API
// @version 1.0
// @description This is the backend API for Animal NGO Universal.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@alngo.org

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig()
	db, e := database.ConnectDatabase(cfg)
	if e != nil {
		log.Fatal(e)
	}
	if e = database.AutoMigrate(db); e != nil {
		log.Fatal(e)
	}
	dh := handlers.NewDepartmentHandler(services.NewDepartmentService(repositories.NewDepartmentRepository(db)))
	ph := handlers.NewPositionHandler(services.NewPositionService(repositories.NewPositionRepository(db)))
	wh := handlers.NewWorkflowHandler(services.NewWorkflowService(repositories.NewWorkflowRepository(db)))
	sh := handlers.NewStepHandler(services.NewStepService(repositories.NewStepRepository(db)))
	eh := handlers.NewEntryHandler(services.NewEntryService(repositories.NewEntryRepository(db)))
	
	otpStore := services.NewOTPStore()
	smsSvc := services.NewSMSService()
	ah := handlers.NewAuthHandler(services.NewAuthService(db, cfg, otpStore, smsSvc))
	
	fh := handlers.NewFormHandler(services.NewFormService(db))
	rh := handlers.NewRuleHandler(services.NewRuleService(db))
	r := gin.Default()
	
	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true, "message": "Animal NGO Universal Backend is running"})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	auth := api.Group("/auth")
	auth.POST("/send-otp/signup", ah.SendSignupOTP)
	auth.POST("/send-otp/login", ah.SendLoginOTP)
	auth.POST("/verify-otp", ah.VerifyOTP)
	
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(cfg))
	protected.GET("/auth/me", ah.Me)

	departments := protected.Group("/departments")
	departments.POST("", dh.Create)
	departments.GET("", dh.List)
	departments.GET("/:id", dh.Get)
	departments.PUT("/:id", dh.Update)
	departments.DELETE("/:id", dh.Delete)
	departments.GET("/:id/positions", ph.ListByDepartment)

	positions := protected.Group("/positions")
	positions.POST("", ph.Create)
	positions.GET("", ph.List)
	positions.GET("/:id", ph.Get)
	positions.PUT("/:id", ph.Update)
	positions.DELETE("/:id", ph.Delete)

	workflows := protected.Group("/workflows")
	workflows.POST("", wh.Create)
	workflows.GET("", wh.List)
	workflows.GET("/:id", wh.Get)
	workflows.PUT("/:id", wh.Update)
	workflows.DELETE("/:id", wh.Delete)
	workflows.POST("/:id/steps", sh.Create)
	workflows.PUT("/:id/steps/reorder", sh.Reorder)
	workflows.POST("/:id/steps/:step_id/submit", eh.Create)

	entries := protected.Group("/entries")
	entries.GET("", eh.List)
	entries.GET("/:id", eh.Get)

	steps := protected.Group("/steps")
	steps.POST("/:id/fields", fh.Create)
	steps.POST("/:id/rules", rh.Create)

	forms := protected.Group("/form-fields")
	forms.GET("/step/:stepID", fh.List)
	forms.DELETE("/:id", fh.Delete)
	log.Printf("Swagger UI available at http://localhost:%s/swagger/index.html\n", cfg.AppPort)
	log.Println("Server starting on port", cfg.AppPort)
	if e = r.Run(":" + cfg.AppPort); e != nil {
		log.Fatal(e)
	}
}
