// routes/api.go

package routes

import (
	"gumuruh-clinic/controllers"
	"gumuruh-clinic/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAPIRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize controllers
	authController := controllers.NewAuthController(db)
	adminController := controllers.NewAdminController(db)
	pasienController := controllers.NewPasienController(db)
	dokterController := controllers.NewDokterController(db)
	obatController := controllers.NewObatController(db)
	kunjunganController := controllers.NewKunjunganController(db)

	// Public routes
	api := router.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.RequireAuth()) // Correct middleware usage
	{
		// Admin routes
		admin := protected.Group("/admin")
		{
			admin.POST("", middleware.RequireAdmin(), adminController.Create)
			admin.GET("", middleware.RequireAdmin(), adminController.List)
			admin.GET("/:id", middleware.RequireAdmin(), adminController.Get)
			admin.PUT("/:id", middleware.RequireAdmin(), adminController.Update)
			admin.DELETE("/:id", middleware.RequireAdmin(), adminController.Delete)
		}

		// Pasien routes
		pasien := protected.Group("/pasien")
		{
			pasien.POST("", pasienController.Create)
			pasien.GET("", middleware.RequireAdmin(), pasienController.List)
			pasien.GET("/:id", pasienController.Get)
			pasien.PUT("/:id", middleware.RequireAdminOrSelf(), pasienController.Update)
			pasien.DELETE("/:id", middleware.RequireAdmin(), pasienController.Delete)
		}

		// Dokter routes
		dokter := protected.Group("/dokter")
		{
			dokter.POST("", middleware.RequireAdmin(), dokterController.Create)
			dokter.GET("", dokterController.List)
			dokter.GET("/:id", dokterController.Get)
			dokter.PUT("/:id", middleware.RequireAdminOrSelf(), dokterController.Update)
			dokter.DELETE("/:id", middleware.RequireAdmin(), dokterController.Delete)
		}

		// Obat routes
		obat := protected.Group("/obat")
		{
			obat.POST("", middleware.RequireAdmin(), obatController.Create)
			obat.GET("", obatController.List)
			obat.GET("/:id", obatController.Get)
			obat.PUT("/:id", middleware.RequireAdmin(), obatController.Update)
			obat.DELETE("/:id", middleware.RequireAdmin(), obatController.Delete)
		}

		// Kunjungan routes
		kunjungan := protected.Group("/kunjungan")
		{
			kunjungan.POST("", middleware.RequireAdmin(), kunjunganController.Create)
			kunjungan.GET("", kunjunganController.List)
			kunjungan.GET("/:id", kunjunganController.Get)
			kunjungan.PUT("/:id", middleware.RequireAdminOrDoctor(), kunjunganController.Update)
			kunjungan.DELETE("/:id", middleware.RequireAdmin(), kunjunganController.Delete)
		}
	}
}
