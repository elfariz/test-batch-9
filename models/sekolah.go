package models

type Sekolah struct {
	Id           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama_Sekolah string `json:"nama_sekolah"`
	Kota_Sekolah string `json:"kota_sekolah"`
}
