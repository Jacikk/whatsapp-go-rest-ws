package categories_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (c controllers) Delete(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON("Invalid category ID")
	}
	if id >= 1 && id <= 3 {
		return ctx.Status(400).JSON("Cannot delete default categories")
	}
	
	err = c.Db.Category.DeleteOneID(id).Exec(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON("Error deleting category")
	}
	ctx.SendStatus(200)
	return nil
}
