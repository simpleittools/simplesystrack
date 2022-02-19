package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/simpleittools/simplesystrack/database"
	"github.com/simpleittools/simplesystrack/routes"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.Conn()

	port := os.Getenv("GOSERVERPORT")

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(port))
}
