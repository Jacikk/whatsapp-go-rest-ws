package routes

import (
	"bd_test/internals/server/controllers/challenge_controller"

	"github.com/gofiber/fiber/v3"
)

func RegisterChallengeRouter(app *fiber.App, controller challenge_controller.Controllers) {
	app.Get("/", controller.Get)
}
