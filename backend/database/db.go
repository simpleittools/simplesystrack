package database

import (
	"fmt"
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
	} else {
		fmt.Println("connected to MySQL")
	}

	DB = conn

	conn.AutoMigrate(&models.Client{}, &models.Project{})
}
