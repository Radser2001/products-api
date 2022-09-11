package main

import (
	"log"

	"github.com/Radser2001/products_api/database"
	"github.com/Radser2001/products_api/routes"
	"github.com/gofiber/fiber/v2"
)

func home(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Product Manager")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/api", home)
	app.Post("/api/addProduct", routes.CreateProduct)
	app.Get("/api/getProducts", routes.GetProducts)
	app.Get("/api/getProduct/:id", routes.GetProduct)
	app.Put("/api/updateProduct/:id", routes.UpdateProduct)
	app.Delete("/api/deleteProduct/:id", routes.DeleteProduct)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
