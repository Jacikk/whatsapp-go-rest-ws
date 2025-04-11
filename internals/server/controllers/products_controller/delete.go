package products_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (c controllers) Delete(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON("Invalid product ID")
	}
	if id >= 1 && id <= 20 {
		return ctx.Status(400).JSON("Cannot delete default products")
	}
	err = c.Db.Product.DeleteOneID(id).Exec(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON("Error deleting product")
	}
	ctx.SendStatus(200)
	return nil
}
