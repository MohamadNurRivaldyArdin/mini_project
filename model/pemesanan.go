package model

import "time"

type Pemesanan struct {
	Id            int       `gorm:"primaryKey" json:"Id"`
	Tgl_Pemesanan time.Time `json:"Tgl_Pemesanan"`
	Tgl_Keluar    time.Time `json:"Tgl_Keluar"`
	Pengunjung_Id int       `json:"Pengunjung_Id"`
	Kost_Id       int       `json:"Kost_Id"`
}

func (Pemesanan) TableName() string {
	return "pemesanan"
}
