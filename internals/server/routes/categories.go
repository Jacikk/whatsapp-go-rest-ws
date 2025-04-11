package routes

import (
	"bd_test/internals/server/controllers/categories_controller"

	"github.com/gofiber/fiber/v3"
)

func RegisterCategoriesRouter(app *fiber.App, controller categories_controller.Controllers) {
	router := app.Group("/categories")

	router.Get("/", controller.List)
	router.Get("/:id", controller.Get)
	router.Post("/", controller.Create)
	router.Patch("/:id", controller.Patch)
	router.Delete("/:id", controller.Delete)
}
