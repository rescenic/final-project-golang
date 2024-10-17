package controllers

import (
	"gumuruh-clinic/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DokterController struct {
	db *gorm.DB
}

func NewDokterController(db *gorm.DB) *DokterController {
	return &DokterController{db: db}
}

func (c *DokterController) Create(ctx *gin.Context) {
	var dokter models.Dokter
	if err := ctx.ShouldBindJSON(&dokter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Create(&dokter).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create dokter"})
		return
	}

	ctx.JSON(http.StatusCreated, dokter)
}

func (c *DokterController) List(ctx *gin.Context) {
	var dokters []models.Dokter
	if err := c.db.Find(&dokters).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch dokters"})
		return
	}

	ctx.JSON(http.StatusOK, dokters)
}

func (c *DokterController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	var dokter models.Dokter

	if err := c.db.First(&dokter, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Dokter not found"})
		return
	}

	ctx.JSON(http.StatusOK, dokter)
}

func (c *DokterController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var dokter models.Dokter

	if err := c.db.First(&dokter, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Dokter not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&dokter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Save(&dokter).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update dokter"})
		return
	}

	ctx.JSON(http.StatusOK, dokter)
}

func (c *DokterController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var dokter models.Dokter

	if err := c.db.First(&dokter, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Dokter not found"})
		return
	}

	if err := c.db.Delete(&dokter).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete dokter"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Dokter deleted successfully"})
}
