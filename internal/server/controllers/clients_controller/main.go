package clients_controller

import (
	"whatsapp_rest_ws/internal/server/services/whatsapp_service"

	"github.com/gofiber/fiber/v3"
	socketio "github.com/googollee/go-socket.io"
)

type controller struct {
	whatsappService whatsapp_service.Service
	io              *socketio.Server
}

type Controller interface {
	GetContacts(ctx fiber.Ctx) error
	GetGroups(ctx fiber.Ctx) error
	Connect(ctx fiber.Ctx) error
	Reconnect(ctx fiber.Ctx) error
	GetPairingCode(ctx fiber.Ctx) error
	Disconnect(ctx fiber.Ctx) error
	SendMessage(ctx fiber.Ctx) error
	IsConnected(ctx fiber.Ctx) error
}

func New(io *socketio.Server) Controller {
	return &controller{
		whatsappService: whatsapp_service.New(),
		io:              io,
	}
}
