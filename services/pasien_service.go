// services/pasien_service.go

package services

import (
	"fmt"
	"gumuruh-clinic/models"

	"gorm.io/gorm"
)

// GenerateNoRM increments the last NoRM and returns the new one without "RM-" prefix
func GenerateNoRM(db *gorm.DB) (string, error) {
	var maxNoRM int
	// Query the maximum NoRM from the database
	if err := db.Model(&models.Pasien{}).Select("COALESCE(MAX(CAST(no_rm AS UNSIGNED)), 0)").Scan(&maxNoRM).Error; err != nil {
		return "", err // Handle error
	}
	// Increment and format the new NoRM
	newNoRM := maxNoRM + 1
	return fmt.Sprintf("%06d", newNoRM), nil // Ensure 6 digits without RM-
}
