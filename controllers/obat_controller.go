// controllers/obat_controllers.go

package controllers

import (
	"gumuruh-clinic/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ObatController struct {
	db *gorm.DB
}

func NewObatController(db *gorm.DB) *ObatController {
	return &ObatController{db: db}
}

func (c *ObatController) Create(ctx *gin.Context) {
	var obat models.Obat
	if err := ctx.ShouldBindJSON(&obat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Create(&obat).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create obat"})
		return
	}

	ctx.JSON(http.StatusCreated, obat)
}

func (c *ObatController) List(ctx *gin.Context) {
	var obats []models.Obat
	if err := c.db.Find(&obats).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch obats"})
		return
	}

	ctx.JSON(http.StatusOK, obats)
}

func (c *ObatController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	var obat models.Obat

	if err := c.db.First(&obat, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Obat not found"})
		return
	}

	ctx.JSON(http.StatusOK, obat)
}

func (c *ObatController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var obat models.Obat

	if err := c.db.First(&obat, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Obat not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&obat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Save(&obat).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update obat"})
		return
	}

	ctx.JSON(http.StatusOK, obat)
}

func (c *ObatController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var obat models.Obat

	if err := c.db.First(&obat, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Obat not found"})
		return
	}

	if err := c.db.Delete(&obat).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete obat"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Obat deleted successfully"})
}
