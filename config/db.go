package config

import (
	"github.com/AvinFajarF/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	ConnectionDB()
	ConnectRPC()
}

var DB *gorm.DB
var err error

func ConnectionDB() {
	dsn := "host=127.0.0.1 user=root password=root dbname=root port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(model.UserModel{})

}
