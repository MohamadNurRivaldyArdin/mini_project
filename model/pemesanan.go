package model

import "time"

type Pemesanan struct {
	Id            int        `gorm:"primaryKey" json:"Id"`
	Tgl_Pemesanan time.Time  `json:"Tgl_Pemesanan"`
	Tgl_Keluar    time.Time  `json:"Tgl_Keluar"`
	Pengunjung_Id int        `json:"_"`
	Kost_Id       int        `json:"_"`
	Pengunjung    Pengunjung `json:"IdPengunjung" gorm:"foreignKey:Pengunjung_Id;references:Id"`
	Kost          Kost       `json:"IdKost" gorm:"foreignKey:Kost_Id;references:Id"`
}

func (Pemesanan) TableName() string {
	return "pemesanan"
}
