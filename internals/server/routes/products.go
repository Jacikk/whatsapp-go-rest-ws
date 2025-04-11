package routes

import (
	"bd_test/internals/server/controllers/products_controller"

	"github.com/gofiber/fiber/v3"
)

func RegisterProductsRouter(app *fiber.App, controller products_controller.Controllers) {
	router := app.Group("/products")

	router.Get("/", controller.List)
	router.Get("/:id", controller.Get)
	router.Post("/", controller.Create)
	router.Patch("/:id", controller.Patch)
	router.Delete("/:id", controller.Delete)
}
