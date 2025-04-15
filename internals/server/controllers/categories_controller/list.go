package categories_controller

import (
	"bd_test/internals/server/ent/category"

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

	catQ := c.Db.Category.Query()
	
	if filters.Name != "" {
		catQ = catQ.Where(category.NameContainsFold(ctx.Query("name")))
	}
	
	count, err := catQ.Count(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}
	
	catQ = catQ.Offset(filters.Offset).Limit(filters.Limit)
	
	categories, err := catQ.All(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"rows":  categories,
		"count": count,
	})
}
