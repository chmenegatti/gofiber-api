package routes

import (
	"github.com/chmenegatti/gofiber-api/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Post("/", controller.CreateUser)
	user.Get("/", controller.GetAllUsers)
	user.Get("/:id", controller.GetUser)
	user.Patch("/:id", controller.UpdateUser)
	user.Delete("/:id", controller.DeleteUser)

	products := app.Group("/products")
	products.Post("/", controller.CreateProduct)
	products.Get("/", controller.GetAllProducts)
	products.Get("/:id", controller.GetProduct)
	products.Patch("/:id", controller.UpdateProduct)
	products.Delete("/:id", controller.DeleteProduct)

}
