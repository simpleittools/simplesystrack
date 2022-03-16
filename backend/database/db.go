package database

import (
	"fmt"
	"github.com/simpleittools/simplesystrack/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Conn() {

	if os.Getenv("DATABASE") == "MYSQL" {
		//	MySQL
		dsn := os.Getenv("MYSQLDSN")
		//fmt.Println(dsn)
		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal("Could not connect to the DB")
		} else {
			fmt.Println("connected to MySQL")
		}

		DB = conn

		conn.AutoMigrate(&models.Client{}, &models.Project{}, &models.Milestone{}, &models.Task{})

	} else if os.Getenv("DATABASE") == "PGSQL" {
		dsn := os.Getenv("PGSQLDSN")
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal("Could not connect to the DB")
		} else {
			fmt.Println("connected to PGSQL")
		}

		DB = conn

		conn.AutoMigrate(&models.Client{}, &models.Project{}, &models.Milestone{}, &models.Task{})
	} else {
		log.Fatal("Database Type not supported")
	}

}
