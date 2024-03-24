package main

import (
	"gochat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化 mysql table
func main() {
	dsn := "root:your_password@tcp(localhost:3306)/ginchat?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect darabase")
	}

	// 迁移 schema
	// db.AutoMigrate(&models.UserBasic{})
	// db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.GroupBasic{})
	db.AutoMigrate(&models.Contact{})

}
