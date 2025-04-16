package categories_controller

import (
	"bd_test/internals/server/ent"

	"github.com/gofiber/fiber/v3"
)

type UpdateCategoryRequest struct {
	Id   int     `json:"id"`
	Name *string `json:"name"`
}

func (c controllers) Patch(ctx fiber.Ctx) error {
	req := new(UpdateCategoryRequest)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	req.Id = fiber.Params[int](ctx, "id")

	if req.Id == 0 {
		return ctx.Status(400).JSON("Id is required")
	}

	if req.Id >= 1 && req.Id <= 3 {
		return ctx.Status(400).JSON("Cannot update default categories")
	}

	tx, err := c.Db.Tx(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	defer tx.Rollback()

	updateQ := tx.Category.UpdateOneID(req.Id)

	if req.Name != nil {
		if *req.Name == "" {
			return ctx.Status(400).JSON("Name if present, cannot be empty")
		}

		updateQ.SetName(*req.Name)
	}

	updated, err := updateQ.Save(ctx.Context())
	if err != nil {
		if ent.IsConstraintError(err) {
			return ctx.Status(400).JSON("Category with this name already exists")
		}

		return ctx.Status(500).JSON(err.Error())
	}

	if err := tx.Commit(); err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	return ctx.Status(200).JSON(updated)
}
