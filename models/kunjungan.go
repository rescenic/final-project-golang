// models/kunjungan.go

package models

import "time"

type Kunjungan struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	IDAdmin          uint      `json:"id_admin" gorm:"not null"`
	IDPasien         uint      `json:"id_pasien" gorm:"not null"`
	IDDokter         uint      `json:"id_dokter" gorm:"not null"`
	IDObat           uint      `json:"id_obat" gorm:"not null"`
	TanggalKunjungan time.Time `json:"tanggal_kunjungan" gorm:"not null"`
	RiwayatPenyakit  string    `json:"riwayat_penyakit"`
	Diagnosa         string    `json:"diagnosa"`
	ResepObat        string    `json:"resep_obat"`
	CreatedTime      time.Time `json:"created_time"`
	CreatedBy        string    `json:"created_by"`
	ModifiedTime     time.Time `json:"modified_time"`
	ModifiedBy       string    `json:"modified_by"`
	Active           bool      `json:"active" gorm:"default:true"`
	// Relationships with proper references
	Admin  Admin  `json:"admin" gorm:"foreignKey:IDAdmin"`
	Pasien Pasien `json:"pasien" gorm:"foreignKey:IDPasien"`
	Dokter Dokter `json:"dokter" gorm:"foreignKey:IDDokter"`
	Obat   Obat   `json:"obat" gorm:"foreignKey:IDObat"`
}

func (Kunjungan) TableName() string {
	return "kunjungan"
}
