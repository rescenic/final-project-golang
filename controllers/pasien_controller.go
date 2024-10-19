// controllers/pasien_controller.go

package controllers

import (
	"fmt"
	"gumuruh-clinic/middleware"
	"gumuruh-clinic/models"
	"gumuruh-clinic/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PasienController struct {
	db *gorm.DB
}

func NewPasienController(db *gorm.DB) *PasienController {
	return &PasienController{db: db}
}

func (c *PasienController) Create(ctx *gin.Context) {
	var pasien models.Pasien
	if err := ctx.ShouldBindJSON(&pasien); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate the next NoRM
	noRM, err := c.GetNextNoRM()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate NoRM"})
		return
	}
	pasien.NoRM = noRM

	// Hash password before saving
	hashedPassword, err := services.HashPassword(pasien.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	pasien.Password = hashedPassword

	if err := c.db.Create(&pasien).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pasien"})
		return
	}

	ctx.JSON(http.StatusCreated, pasien)
}

func (c *PasienController) List(ctx *gin.Context) {
	var pasiens []models.Pasien
	if err := c.db.Find(&pasiens).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pasiens"})
		return
	}

	ctx.JSON(http.StatusOK, pasiens)
}

func (c *PasienController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	var pasien models.Pasien

	if err := c.db.First(&pasien, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Pasien not found"})
		return
	}

	ctx.JSON(http.StatusOK, pasien)
}

func (c *PasienController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var existingPasien models.Pasien

	// Find the existing pasien by ID
	if err := c.db.First(&existingPasien, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Pasien not found"})
		return
	}

	// Bind the incoming JSON but keep the existing ID
	var updatedPasien models.Pasien
	if err := ctx.ShouldBindJSON(&updatedPasien); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the existing ID is preserved
	updatedPasien.ID = existingPasien.ID

	// If password is updated, hash the new password
	if updatedPasien.Password != "" && updatedPasien.Password != existingPasien.Password {
		hashedPassword, err := services.HashPassword(updatedPasien.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		updatedPasien.Password = hashedPassword
	} else {
		updatedPasien.Password = existingPasien.Password
	}

	// Keep NoRM consistent
	updatedPasien.NoRM = existingPasien.NoRM

	// Set active to true if not provided
	if updatedPasien.Active == false {
		updatedPasien.Active = true
	}

	// Get the user from the token
	userToken, claims, err := middleware.GetUserFromToken(ctx)
	if err != nil || !userToken.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Set modified_by to the logged-in user
	if name, ok := claims["nama_lengkap"].(string); ok {
		updatedPasien.ModifiedBy = name
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to extract user information from token"})
		return
	}

	// Set modified_time to now
	updatedPasien.ModifiedTime = time.Now()

	// Keep created_time and created_by consistent
	updatedPasien.CreatedTime = existingPasien.CreatedTime
	updatedPasien.CreatedBy = existingPasien.CreatedBy

	// Save the updated pasien
	if err := c.db.Model(&existingPasien).Updates(updatedPasien).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pasien"})
		return
	}

	ctx.JSON(http.StatusOK, updatedPasien)
}

func (c *PasienController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var pasien models.Pasien

	if err := c.db.First(&pasien, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Pasien not found"})
		return
	}

	if err := c.db.Delete(&pasien).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pasien"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pasien deleted successfully"})
}

// GetNextNoRM generates the next NoRM value without the "RM-" prefix
func (c *PasienController) GetNextNoRM() (string, error) {
	var lastPasien models.Pasien

	// Find the latest Pasien based on the highest NoRM
	if err := c.db.Order("no_rm desc").First(&lastPasien).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// If no records are found, return the default NoRM starting point
			return "000001", nil // Starting NoRM
		}
		// Log the error for debugging
		fmt.Println("Error retrieving last pasien:", err)
		return "", err // Return error for other issues
	}

	// Ensure lastPasien.NoRM can be converted to an integer
	lastNoRM, err := strconv.Atoi(lastPasien.NoRM) // No need to skip prefix, full NoRM is numeric
	if err != nil {
		fmt.Println("Error converting NoRM to int:", err)
		return "", err
	}

	// Increment the last NoRM and format it as a 6-digit number
	nextNoRM := fmt.Sprintf("%06d", lastNoRM+1) // Return just the 6 digits
	return nextNoRM, nil
}
