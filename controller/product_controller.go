package controller

import (
	"PCS-Web-Cart/db"
	"PCS-Web-Cart/models"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	db := db.ConnectDB()

	rows, err := db.Query("SELECT * FROM products order by ProductID")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	defer db.Close()

	result := models.Products{}

	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description); err != nil {
			return err
		}

		result.Products = append(result.Products, product)
	}

	return c.JSON(result)
}

func AddProduct(c *fiber.Ctx) error {
	db := db.ConnectDB()

	u := new(models.Product)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	rows, err := db.Query("INSERT INTO products (ProductID, Name, Description) VALUES (?,?)", u.ID, u.Name, u.Description)
	if err != nil {
		return err
	}

	defer rows.Close()
	defer db.Close()

	return c.JSON(u)
}

func DeleteProduct(c *fiber.Ctx) error {
	db := db.ConnectDB()

	u := new(models.Product)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	rows, err := db.Query("DELETE FROM products WHERE ProductID=?", u.ID)
	if err != nil {
		return err
	}

	defer rows.Close()
	defer db.Close()

	return c.JSON("Data Deleted")
}
