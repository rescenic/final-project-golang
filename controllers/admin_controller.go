// controllers/admin_controller.go

package controllers

import (
	"gumuruh-clinic/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminController struct {
	db *gorm.DB
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{db: db}
}

func (c *AdminController) Create(ctx *gin.Context) {
	var admin models.Admin
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Create(&admin).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
		return
	}

	ctx.JSON(http.StatusCreated, admin)
}

func (c *AdminController) List(ctx *gin.Context) {
	var admins []models.Admin
	if err := c.db.Find(&admins).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch admins"})
		return
	}

	ctx.JSON(http.StatusOK, admins)
}

func (c *AdminController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	var admin models.Admin

	if err := c.db.First(&admin, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	ctx.JSON(http.StatusOK, admin)
}

func (c *AdminController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var admin models.Admin

	if err := c.db.First(&admin, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Save(&admin).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin"})
		return
	}

	ctx.JSON(http.StatusOK, admin)
}

func (c *AdminController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var admin models.Admin

	if err := c.db.First(&admin, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	if err := c.db.Delete(&admin).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete admin"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}
