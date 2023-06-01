package Server

import (
	"AIDS_Trigger/Message"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/aids_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error("Error connecting to db")
	}

	err = db.AutoMigrate(&Message.DetectedMessage{})

	if err != nil {
		return nil
	}

	return db
}
