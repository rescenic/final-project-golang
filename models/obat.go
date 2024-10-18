// models/obat.go

package models

import "time"

type Obat struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	NamaObat     string    `json:"nama_obat" gorm:"not null"`
	JenisObat    string    `json:"jenis_obat" gorm:"not null"`
	CreatedTime  time.Time `json:"created_time"`
	CreatedBy    string    `json:"created_by"`
	ModifiedTime time.Time `json:"modified_time"`
	ModifiedBy   string    `json:"modified_by"`
	Active       bool      `json:"active" gorm:"default:true"`
	// Relationships
	Kunjungans []Kunjungan `json:"kunjungans,omitempty" gorm:"foreignKey:IDObat"`
}

func (Obat) TableName() string {
	return "obat"
}