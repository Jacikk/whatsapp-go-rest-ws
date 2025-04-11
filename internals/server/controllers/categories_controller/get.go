package categories_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (c controllers) Get(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON("Invalid category ID")
	}
	category, err := c.Db.Category.Get(ctx.Context(), id)
	if err != nil {
		return ctx.Status(500).JSON("Error getting category")
	}

	return ctx.Status(200).JSON(category)
}
