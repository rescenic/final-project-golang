// models/admin.go

package models

import (
	"time"
)

type Admin struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	IDKtp        string    `json:"id_ktp" gorm:"not null;unique"`
	NamaLengkap  string    `json:"nama_lengkap" gorm:"not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	Password     string    `json:"password,omitempty" gorm:"not null"`
	CreatedTime  time.Time `json:"created_time"`
	CreatedBy    string    `json:"created_by"`
	ModifiedTime time.Time `json:"modified_time"`
	ModifiedBy   string    `json:"modified_by"`
	Active       bool      `json:"active" gorm:"default:true"`
	// Relationships
	Kunjungans []Kunjungan `json:"kunjungans,omitempty" gorm:"foreignKey:IDAdmin"`
}
