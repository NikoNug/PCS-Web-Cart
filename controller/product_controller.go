package controller

import (
	"PCS-Web-Cart/db"
	"PCS-Web-Cart/models"

	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {
	db := db.ConnectDB()

	id := c.Params("id")
	var product models.Product

	rows, err := db.Query("SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	defer db.Close()

	for rows.Next() {
		rows.Scan(&product.ID, &product.Name, &product.Desc)
	}

	return c.JSON(product)
}
