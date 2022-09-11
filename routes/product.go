package routes

import (
	"errors"

	"github.com/Radser2001/products_api/database"
	"github.com/Radser2001/products_api/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

//create products
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

//get all products
func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)
	responseProducts := []Product{}

	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

//find product by id
func FindProductHelper(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}

	if err := FindProductHelper(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

//Update product
func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}

	if err := FindProductHelper(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdatedProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)

}

//delete product
func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}
	if err := FindProductHelper(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted the Product")
}
