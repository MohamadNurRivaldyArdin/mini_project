package model

type Admin struct {
	Id         int    `gorm:"primaryKey" json:"Id"`
	Nama       string `json:"Nama"`
	Email      string `json:"Email"`
	Password   int    `json:"Password"`
	Kost_Id    int    `json:"Kost_Id"`
	Pemilik_Id int    `json:"Pemilik_Id "`
}

// type GuestResponse struct {
// 	Id    int    `json: "id" form: "id"`
// 	Email string `json: "email" form: "email"`
// 	Name  string `json: "name" form: "name"`
// 	Token string `json: "token" form: "token"`
// }

func (Admin) TableName() string {
	return "admin"
}
