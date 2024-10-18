// main.go

package main

import (
	"fmt"
	"gumuruh-clinic/config"
	migrate "gumuruh-clinic/migrations"
	"gumuruh-clinic/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migration
	migrate.Migrate(db)

	// Create Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Load HTML templates and static files
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Setup routes
	routes.SetupAPIRoutes(router, db)
	setupViewRoutes(router)

	// Start server
	port := ":8090"
	log.Printf("Server is starting on port %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupViewRoutes(router *gin.Engine) {
	// Home page
	router.GET("/", func(c *gin.Context) {
		c.File("./templates/index.html")
	})

	// Auth routes
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	// Dashboard routes
	router.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})

	// Master data routes
	router.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", nil)
	})

	router.GET("/pasien", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pasien/index.html", nil)
	})

	router.GET("/dokter", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dokter/index.html", nil)
	})

	router.GET("/obat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "obat/index.html", nil)
	})

	// Kunjungan routes
	router.GET("/kunjungan", func(c *gin.Context) {
		c.HTML(http.StatusOK, "kunjungan/index.html", nil)
	})
}
