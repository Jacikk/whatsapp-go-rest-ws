package products_controller

import (
	"bd_test/internals/server/ent"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (c controllers) Get(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON("Invalid product ID")
	}
	category, err := c.Db.Product.Get(ctx.Context(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.Status(404).JSON("Product not found")
		}
		return ctx.Status(500).JSON("Error getting product")
	}

	return ctx.Status(200).JSON(category)
}
