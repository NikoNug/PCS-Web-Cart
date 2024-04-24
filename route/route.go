package route

import (
	"PCS-Web-Cart/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(c *fiber.App) {
	// Home
	// c.Get("/", controller)

	product := c.Group("/product")
	product.Get("/", controller.GetProduct)
}
