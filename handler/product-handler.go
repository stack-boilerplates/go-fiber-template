package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/vserpa/go-fiber-template/database"
	"github.com/vserpa/go-fiber-template/model"
)

func GetAllProducts(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT name, description, category, amount FROM products order by name")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	result := model.Products{}

	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result.Products = append(result.Products, product)
	}

	if err := rows.Err(); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"product": result,
		"message": "All product returned successfully",
	})
}

func GetSingleProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := model.Product{}

	row, err := database.DB.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&id, &product.Amount, &product.Name, &product.Description, &product.Category)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"product": product,
		"message": "Product returned successfully",
	})
}

func CreateProduct(c *fiber.Ctx) error {
	p := new(model.Product)

	if err := c.BodyParser(p); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	_, err := database.DB.Query("INSERT INTO products (name, description, category, amount) VALUES ($1, $2, $3, $4)", p.Name, p.Description, p.Category, p.Amount)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Product created successfully",
		"product": p,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Query("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Product deleted successfully: " + id,
	})
}
