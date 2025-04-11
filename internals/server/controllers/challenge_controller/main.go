package challenge_controller

import (
	"github.com/gofiber/fiber/v3"
)

type controllers struct{}

type Controllers interface {
	Get(c fiber.Ctx) error
}

func New() Controllers {
	return &controllers{}
}
