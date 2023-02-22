package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Donation struct {
	gorm.Model
	Amount int  `json:"amount"`
	UserId uint `json:"user_id"`
}

var DB *gorm.DB

func InitDatabase() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//数据库迁移
	DB.AutoMigrate(&User{}, &Donation{})
	return nil
}
