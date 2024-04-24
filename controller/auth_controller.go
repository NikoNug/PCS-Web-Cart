package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetID(c *fiber.Ctx) error {
	fmt.Println("OK")

	return c.JSON("OK")
}
