// routes/api.go

package routes

import (
	"net/http"

	"gumuruh-clinic/controllers"
	"gumuruh-clinic/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServerInfo struct {
	Message string      `json:"message"`
	Owner   string      `json:"owner"`
	Routes  []RouteInfo `json:"routes"`
}

type RouteInfo struct {
	Path    string   `json:"path"`
	Methods []string `json:"methods"`
}

func SetupAPIRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize controllers
	authController := controllers.NewAuthController(db)
	adminController := controllers.NewAdminController(db)
	pasienController := controllers.NewPasienController(db)
	dokterController := controllers.NewDokterController(db)
	obatController := controllers.NewObatController(db)
	kunjunganController := controllers.NewKunjunganController(db)

	routes := []RouteInfo{
		{Path: "/api/register", Methods: []string{"POST"}},
		{Path: "/api/login", Methods: []string{"POST"}},
		{Path: "/api/admin", Methods: []string{"POST", "GET"}},
		{Path: "/api/admin/:id", Methods: []string{"GET", "PUT", "DELETE"}},
		{Path: "/api/pasien", Methods: []string{"POST", "GET"}},
		{Path: "/api/pasien/:id", Methods: []string{"GET", "PUT", "DELETE"}},
		{Path: "/api/dokter", Methods: []string{"POST", "GET"}},
		{Path: "/api/dokter/:id", Methods: []string{"GET", "PUT", "DELETE"}},
		{Path: "/api/obat", Methods: []string{"POST", "GET"}},
		{Path: "/api/obat/:id", Methods: []string{"GET", "PUT", "DELETE"}},
		{Path: "/api/kunjungan", Methods: []string{"POST", "GET"}},
		{Path: "/api/kunjungan/:id", Methods: []string{"GET", "PUT", "DELETE"}},
	}

	// Public routes
	api := router.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)

		api.GET("/", func(c *gin.Context) {
			info := ServerInfo{
				Message: "Server API Klinik Gumuruh",
				Owner:   "Muhammad Ridwan Hakim, S.T., CPITA",
				Routes:  routes,
			}
			c.JSON(http.StatusOK, info)
		})
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.RequireAuth())
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
