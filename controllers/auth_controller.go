// controllers/auth_controller.go

package controllers

import (
	"fmt"
	"gumuruh-clinic/config"
	"gumuruh-clinic/models"
	"gumuruh-clinic/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	db          *gorm.DB
	authService *services.AuthService
}

func NewAuthController(db *gorm.DB) *AuthController {
	// Load JWT secret from config (which gets it from .env)
	cfg := config.LoadConfig()
	return &AuthController{
		db:          db,
		authService: services.NewAuthService(db, cfg.JWTSecret), // Use JWTSecret from .env
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.authService.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *AuthController) Register(ctx *gin.Context) {
	var pasien models.Pasien
	if err := ctx.ShouldBindJSON(&pasien); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password before saving
	hashedPassword, err := services.HashPassword(pasien.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}
	pasien.Password = hashedPassword

	// Save the new pasien to the database
	if err := c.db.Create(&pasien).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Set NoRM to an incremented value based on the ID
	pasien.NoRM = fmt.Sprintf("%06d", pasien.ID)

	// Update the record with the generated NoRM
	if err := c.db.Save(&pasien).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update NoRM"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}
