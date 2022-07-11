package model

import "time"

type Pemesanan struct {
	ID_Pemesanan  int       `gorm:"primaryKey" json:"ID_Pemesanan"`
	Tgl_Pemesanan time.Time `json:"Tgl_Pemesanan"`
	Tgl_Keluar    time.Time `json:"Tgl_Keluar"`
	PENGUNJUNG_ID int       `json:"PENGUNJUNG_ID"`
	KOST_ID_Kost  int       `json:"KOST_ID_Kost"`
}

func (Pemesanan) TableName() string {
	return "pemesanan"
}
