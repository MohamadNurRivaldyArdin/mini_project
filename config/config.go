package config

import (
	"mini_project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:30042001@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&model.Admin{})
}

func InitDBTest() {
	dsn := "root:30042001@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrateTest() {
	DB.Migrator().DropTable(&model.Admin{})
	DB.AutoMigrate(&model.Admin{})
}
