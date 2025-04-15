package products_controller

import (
	"bd_test/internals/server/ent/product"

	"github.com/gofiber/fiber/v3"
)

type Filters struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Name   string `json:"name"`
}

func (c controllers) List(ctx fiber.Ctx) error {

	filters := Filters{
		Offset: fiber.Query(ctx, "offset", 0),
		Limit:  fiber.Query(ctx, "limit", 10),
		Name:   fiber.Query(ctx, "name", ""),
	}
	if filters.Offset < 0 {
		return ctx.Status(400).JSON("Offset must be equals or greater than 0")
	}

	if filters.Limit <= 0 || filters.Limit > 20 {
		return ctx.Status(400).JSON("Limit must be greater than 0 and less than or equal to 20")
	}

	prodQ := c.Db.Product.Query()

	if filters.Name != "" {
		prodQ = prodQ.Where(product.NameContainsFold(filters.Name))
	}

	count, err := prodQ.Count(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	prodQ = prodQ.Offset(filters.Offset).Limit(filters.Limit)

	products, err := prodQ.All(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"rows":  products,
		"count": count,
	})
}
