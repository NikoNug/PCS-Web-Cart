package route

import (
	"PCS-Web-Cart/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(c *fiber.App) {
	// Home
	c.Get("/", controller.GetHelloWorld)

	product := c.Group("/product")
	product.Get("/products", controller.GetProducts)
	product.Post("/products", controller.AddProduct)
	product.Delete("/products", controller.DeleteProduct)
}
