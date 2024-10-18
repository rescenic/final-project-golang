// controllers/auth_controller.go

package controllers

import (
	"gumuruh-clinic/config"
	"gumuruh-clinic/models"
	"gumuruh-clinic/services"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	db          *gorm.DB
	authService *services.AuthService
}

func NewAuthController(db *gorm.DB) *AuthController {
	cfg := config.LoadConfig()
	return &AuthController{
		db:          db,
		authService: services.NewAuthService(db, cfg.JWTSecret),
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var admin models.Admin
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Validate required fields
	if admin.IDKtp == "" || admin.NamaLengkap == "" || admin.Email == "" || admin.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "All fields (id_ktp, nama_lengkap, email, password) are required"})
		return
	}

	// Validate ID KTP length
	if len(admin.IDKtp) != 16 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID KTP must be 16 characters long"})
		return
	}

	// Check if email already exists
	var existingAdmin models.Admin
	if err := c.db.Where("email = ?", admin.Email).First(&existingAdmin).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Check if ID KTP already exists
	if err := c.db.Where("id_ktp = ?", admin.IDKtp).First(&existingAdmin).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "ID KTP already registered"})
		return
	}

	// Hash password
	hashedPassword, err := services.HashPassword(admin.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	// Prepare admin data
	admin.Password = hashedPassword
	admin.CreatedTime = time.Now()
	admin.ModifiedTime = time.Now()
	admin.CreatedBy = "SYSTEM"
	admin.ModifiedBy = "SYSTEM"
	admin.Active = true

	// Create admin
	if err := c.db.Create(&admin).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "unique") {
			if strings.Contains(strings.ToLower(err.Error()), "email") {
				ctx.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
			} else if strings.Contains(strings.ToLower(err.Error()), "id_ktp") {
				ctx.JSON(http.StatusConflict, gin.H{"error": "ID KTP already registered"})
			} else {
				ctx.JSON(http.StatusConflict, gin.H{"error": "Duplicate entry detected"})
			}
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin: " + err.Error()})
		return
	}

	// Remove password from response
	admin.Password = ""
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Admin created successfully",
		"data":    admin,
	})
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
