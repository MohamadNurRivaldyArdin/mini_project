package model

type Pemilik struct {
	Id            int    `gorm:"primaryKey" json:"Id"`
	Nama          string `json:"Nama"`
	Alamat        string `json:"Alamat"`
	Telepon       int    `json:"Telepon"`
	Email         string `json:"Email"`
	Username      string `json:"Username"`
	Password      int    `json:"Password"`
	Jenis_Kelamin string `json:"Jenis_Kelamin"`
}

func (Pemilik) TableName() string {
	return "pemilik"
}
