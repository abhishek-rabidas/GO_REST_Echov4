package Server

import (
	"AIDS_Trigger/Message"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type DatabaseInfo struct {
	Username string
	Password string
	Host     string
	Database string
}

func InitializeDB() *gorm.DB {
	//dsn := "admin:password@tcp(aids-db.cbo2bmr3r0wv.ap-south-1.rds.amazonaws.com:3306)/aids_db?parseTime=true"
	var databaseInfo DatabaseInfo = DatabaseInfo{}
	dbConfig, err := os.ReadFile("./db.json")
	err = json.Unmarshal(dbConfig, &databaseInfo)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseInfo.Username, databaseInfo.Password, databaseInfo.Host, databaseInfo.Database)
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
