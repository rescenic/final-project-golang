// controllers/pasien_controller.go

package controllers

import (
	"gumuruh-clinic/models"
	"gumuruh-clinic/services"
	"net/http"

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
