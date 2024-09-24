package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/vserpa/go-fiber-template/database"
	"github.com/vserpa/go-fiber-template/dto"
	"github.com/vserpa/go-fiber-template/model"
)

func GetAllProducts(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT name, description, category, amount FROM products order by name")
	if err != nil {
		return c.Status(500).JSON(dto.ResponseFailed(err))
	}
	defer rows.Close()

	result := model.Products{}

	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount)
		if err != nil {
			return c.Status(500).JSON(dto.ResponseFailed(err))
		}

		result.Products = append(result.Products, product)
	}

	if err := rows.Err(); err != nil {
		return c.Status(500).JSON(dto.ResponseFailed(err))
	}

	return c.Status(200).JSON(dto.ResponseSuccess(result.Products))
}

func GetSingleProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := model.Product{}

	row, err := database.DB.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(dto.ResponseFailed(err))
	}
	defer row.Close()

	if !row.Next() {
		return c.Status(404).JSON(dto.ResponseFailed(fmt.Errorf("product not found")))
	}

	err = row.Scan(&id, &product.Amount, &product.Name, &product.Description, &product.Category)
	if err != nil {
		return c.Status(500).JSON(dto.ResponseFailed(err))
	}

	return c.Status(200).JSON(dto.ResponseSuccess(product))
}

func CreateProduct(c *fiber.Ctx) error {
	p := new(model.Product)

	if err := c.BodyParser(p); err != nil {
		return c.Status(400).JSON(dto.ResponseFailed(err))
	}

	_, err := database.DB.Query("INSERT INTO products (name, description, category, amount) VALUES ($1, $2, $3, $4)", p.Name, p.Description, p.Category, p.Amount)
	if err != nil {
		return c.Status(500).JSON(dto.ResponseFailed(err))
	}

	return c.Status(201).JSON(dto.ResponseSuccess(p))
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := database.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(dto.ResponseFailed(err))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(500).JSON(dto.ResponseFailed(err))
	}

	if rowsAffected == 0 {
		return c.Status(404).JSON(dto.ResponseFailed(fmt.Errorf("product not found with id: %s", id)))
	}

	return c.Status(200).JSON(dto.ResponseSuccess(id))
}
