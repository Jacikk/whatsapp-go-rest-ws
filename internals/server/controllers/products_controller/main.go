package products_controller

import (
	"bd_test/internals/server/ent"

	"github.com/gofiber/fiber/v3"
)

type controllers struct {
	Db *ent.Client
}

type Controllers interface {
	Get(c fiber.Ctx) error
	Create(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
	List(c fiber.Ctx) error
	Patch(c fiber.Ctx) error
}

func New(db *ent.Client) Controllers {
	return &controllers{
		Db: db,
	}
}
