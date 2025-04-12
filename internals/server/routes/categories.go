package routes

import (
	"bd_test/internals/server/controllers/categories_controller"
	"bd_test/internals/server/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterCategoriesRouter(app *fiber.App, controller categories_controller.Controllers) {
	mdlwares := middlewares.New(controller.GetDbInstance())
	router := app.Group("/categories")

	router.Use(mdlwares.AccessKeyMiddleware())

	router.Get("/", controller.List)
	router.Get("/:id", controller.Get)
	router.Post("/", controller.Create)
	router.Patch("/:id", controller.Patch)
	router.Delete("/:id", controller.Delete)
}
