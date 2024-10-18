// models/pasien.go

package models

import "time"

type Pasien struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	IDKtp         string    `json:"id_ktp" gorm:"not null;unique"`
	NoRM          string    `json:"no_rm" gorm:"unique;not null"`
	NamaLengkap   string    `json:"nama_lengkap" gorm:"not null"`
	Email         string    `json:"email" gorm:"unique;not null"`
	Password      string    `json:"password,omitempty" gorm:"not null"`
	TanggalLahir  time.Time `json:"tanggal_lahir"`
	GolonganDarah string    `json:"golongan_darah"`
	CreatedTime   time.Time `json:"created_time"`
	CreatedBy     string    `json:"created_by"`
	ModifiedTime  time.Time `json:"modified_time"`
	ModifiedBy    string    `json:"modified_by"`
	Active        bool      `json:"active" gorm:"default:true"`
	// Relationships
	Kunjungans []Kunjungan `json:"kunjungans,omitempty" gorm:"foreignKey:IDPasien"`
}

func (Pasien) TableName() string {
	return "pasien"
}
