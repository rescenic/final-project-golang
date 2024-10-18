// controllers/pasien_controller.go

package controllers

import (
	"fmt"
	"gumuruh-clinic/models"
	"gumuruh-clinic/services"
	"net/http"
	"strconv"

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
	var pasien models.Pasien

	if err := c.db.First(&pasien, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Pasien not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&pasien); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Save(&pasien).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pasien"})
		return
	}

	ctx.JSON(http.StatusOK, pasien)
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

// GetNextNoRM generates the next NoRM value
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
	lastNoRM, err := strconv.Atoi(lastPasien.NoRM[4:]) // Skip "RM-" prefix
	if err != nil {
		fmt.Println("Error converting NoRM to int:", err)
		return "", err
	}

	// Increment the last NoRM and format it
	nextNoRM := fmt.Sprintf("RM-%06d", lastNoRM+1) // Ensure it has leading zeros
	return nextNoRM, nil
}
