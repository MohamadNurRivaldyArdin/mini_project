package model

type Pemilik struct {
	Id            int    `gorm:"primaryKey" json:"Id"`
	Nama          string `json:"Nama"`
	Alamat        string `json:"Alamat"`
	Telepon       int    `json:"Telepon"`
	Email         string `json:"Email"`
	Username      string `json:"Username"`
	Password      string `json:"Password"`
	Jenis_Kelamin string `json:"Jenis_Kelamin"`
}

type PemilikResponse struct {
	ID    int    `json: "ID" form: "ID"`
	Email string `json: "email" form: "Email"`
	Nama  string `json: "name" form: "Nama"`
	Token string `json: "token" form: "Token"`
}

func (Pemilik) TableName() string {
	return "pemilik"
}
