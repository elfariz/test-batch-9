package models

type Siswa struct {
	Id        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Nisn      int     `json:"nisn"`
	Nama      string  `json:"nama"`
	Alamat    string  `json:"alamat"`
	SekolahId uint    `json:"sekolah_id"`
	Sekolah   Sekolah `json:"sekolah"`
}
