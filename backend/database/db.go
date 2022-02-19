package database

import (
	"github.com/simpleittools/simplesystrack/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Conn() {

	//	MySQL
	dsn := os.Getenv("DSN")
	//fmt.Println(dsn)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the DB")
	}

	DB = conn

	conn.AutoMigrate(&models.Client{})
}
