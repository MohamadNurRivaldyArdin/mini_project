package model

type Admin struct {
	ID           int    `gorm:"primaryKey" json:"ID"`
	User_Name    string `json:"User_Name"`
	Password     int    `json:"Password"`
	KOST_ID_Kost int    `json:"KOST_ID_Kost"`
	PEMILIK_ID   int    `json:"PEMILIK_ID "`
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
