package controllers

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func Index(c *fiber.Ctx) error {
	frontend := os.Getenv("FRONTEND")
	return c.Redirect(frontend)
}
