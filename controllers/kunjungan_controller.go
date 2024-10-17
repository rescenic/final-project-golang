// controllers/kunjungan_controller.go

package controllers

import (
	"gumuruh-clinic/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type KunjunganController struct {
	db *gorm.DB
}

func NewKunjunganController(db *gorm.DB) *KunjunganController {
	return &KunjunganController{db: db}
}

func (c *KunjunganController) Create(ctx *gin.Context) {
	var kunjungan models.Kunjungan
	if err := ctx.ShouldBindJSON(&kunjungan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kunjungan.CreatedTime = time.Now()
	if err := c.db.Create(&kunjungan).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create kunjungan"})
		return
	}

	ctx.JSON(http.StatusCreated, kunjungan)
}

func (c *KunjunganController) List(ctx *gin.Context) {
	var kunjungan []models.Kunjungan
	if err := c.db.Preload("Pasien").Preload("Dokter").Preload("Obat").Find(&kunjungan).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch kunjungan"})
		return
	}

	ctx.JSON(http.StatusOK, kunjungan)
}

func (c *KunjunganController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	var kunjungan models.Kunjungan

	if err := c.db.Preload("Pasien").Preload("Dokter").Preload("Obat").First(&kunjungan, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Kunjungan not found"})
		return
	}

	ctx.JSON(http.StatusOK, kunjungan)
}

func (c *KunjunganController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var kunjungan models.Kunjungan

	if err := c.db.First(&kunjungan, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Kunjungan not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&kunjungan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kunjungan.ModifiedTime = time.Now()
	if err := c.db.Save(&kunjungan).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update kunjungan"})
		return
	}

	ctx.JSON(http.StatusOK, kunjungan)
}

func (c *KunjunganController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var kunjungan models.Kunjungan

	if err := c.db.First(&kunjungan, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Kunjungan not found"})
		return
	}

	if err := c.db.Delete(&kunjungan).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete kunjungan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Kunjungan deleted successfully"})
}
