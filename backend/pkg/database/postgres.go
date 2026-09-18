package database

import (
	"animalngo-universal-backend/config"
	"animalngo-universal-backend/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)
	db, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		return nil, e
	}
	sqlDB, e := db.DB()
	if e != nil {
		return nil, e
	}
	if e = sqlDB.Ping(); e != nil {
		return nil, e
	}
	log.Println("PostgreSQL connected")
	return db, nil
}
func AutoMigrate(db *gorm.DB) error {
	// Drop legacy columns that were removed from the User model
	if db.Migrator().HasColumn(&models.User{}, "email") {
		db.Migrator().DropColumn(&models.User{}, "email")
	}
	if db.Migrator().HasColumn(&models.User{}, "password") {
		db.Migrator().DropColumn(&models.User{}, "password")
	}
	
	return db.AutoMigrate(&models.Department{}, &models.Position{}, &models.User{}, &models.Workflow{}, &models.WorkflowStep{}, &models.FormField{}, &models.StepRule{}, &models.DynamicEntry{}, &models.Patient{})
}
