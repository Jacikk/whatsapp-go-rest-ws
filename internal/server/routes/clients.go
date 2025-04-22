package routes

import (
	"whatsapp_rest_ws/internal/server/controllers/clients_controller"

	"github.com/gofiber/fiber/v3"
)

func RegisterClientsRouter(app *fiber.App, controller clients_controller.Controller) {
	router := app.Group("/clients")

	router.Get("/:id/contacts", controller.GetContacts)
	router.Get("/:id/groups", controller.GetGroups)
	router.Post("/:id/connect", controller.Connect)
	router.Post("/:id/reconnect", controller.Reconnect)
	router.Get("/:id/pairing_code", controller.GetPairingCode)
	router.Post("/:id/disconnect", controller.Disconnect)
	router.Post("/:id/send_message", controller.SendMessage)
	router.Get("/:id/is_connected", controller.IsConnected)
}
