package model

type Kost struct {
	ID_Kost         int    `gorm:"primaryKey" json:"ID_Kost"`
	Nama_Kost       string `json:"Nama_Kost"`
	Jenis_Kost      string `json:"Jenis_Kost"`
	Fasilitas       string `json:"Fasilitas"`
	Harga_per_bulan int    `json:"Harga_per_bulan"`
	Alamat_Kost     string `json:"Alamat_Kost"`
	Kota            string `json:"Kota"`
	Stok            int    `json:"Stok"`
	PEMILIK_ID      int    `json:"PEMILIK_ID"`
}

func (Kost) TableName() string {
	return "kost"
}
