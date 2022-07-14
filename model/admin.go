package model

type Admin struct {
	Id         int    `gorm:"primaryKey" json:"Id"`
	Nama       string `json:"Nama"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	Kost_Id    int    `json:"_"`
	Pemilik_Id int    `json:"_"`
	// Kost       Kost    `json:"Kost" gorm:"foreignKey:Kost_Id_;references:Id"`
	// Pemilik    Pemilik `json:"Pemilik" gorm:"foreignKey:Pemilik_Id;references:Id"`
}

type AdminResponse struct {
	ID    int    `json: "ID" form: "ID"`
	Email string `json: "email" form: "Email"`
	Nama  string `json: "name" form: "Nama"`
	Token string `json: "token" form: "Token"`
}

func (Admin) TableName() string {
	return "admin"
}
