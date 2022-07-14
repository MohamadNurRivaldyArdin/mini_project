package model

type Kost struct {
	Id              int     `gorm:"primaryKey" json:"Id"`
	Nama_Kost       string  `json:"Nama_Kost"`
	Jenis_Kost      string  `json:"Jenis_Kost"`
	Fasilitas       string  `json:"Fasilitas"`
	Harga_per_bulan int     `json:"Harga_per_bulan"`
	Alamat_Kost     string  `json:"Alamat_Kost"`
	Kota            string  `json:"Kota"`
	Stok            int     `json:"Stok"`
	Pemilik_Id      int     `json:"_"`
	Pemilik         Pemilik `json:"IdPemilik" gorm:"foreignKey:Pemilik_Id;references:Id"`
}

func (Kost) TableName() string {
	return "kost"
}
