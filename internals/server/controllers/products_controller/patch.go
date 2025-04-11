package products_controller

import (
	"bd_test/internals/server/ent"

	"github.com/gofiber/fiber/v3"
)

type UpdateProductRequest struct {
	Id         int     `json:"id"`
	Name       *string `json:"name"`
	CategoryId *int    `json:"category_id"`
}

func (c controllers) Patch(ctx fiber.Ctx) error {
	req := new(UpdateProductRequest)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	req.Id = fiber.Params[int](ctx, "id")

	if req.Id == 0 {
		return ctx.Status(400).JSON("Id is required")
	}

	if req.Id >= 1 && req.Id <= 3 {
		return ctx.Status(400).JSON("Cannot update default products")
	}

	tx, err := c.Db.Tx(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	defer tx.Rollback()

	updateQ := tx.Product.UpdateOneID(req.Id)

	if req.Name != nil {
		if *req.Name == "" {
			return ctx.Status(400).JSON("Name if present, cannot be empty")
		}

		updateQ.SetName(*req.Name)
	}

	if req.CategoryId != nil {
		if *req.CategoryId == 0 {
			return ctx.Status(400).JSON("CategoryId if present, cannot be empty")
		}
		if _, err := tx.Category.Get(ctx.Context(), *req.CategoryId); err != nil {
			if ent.IsNotFound(err) {
				return ctx.Status(400).JSON("Category with this id does not exist")
			}
			return ctx.Status(500).JSON(err.Error())
		}
		updateQ.SetCategoryID(*req.CategoryId)
	}

	updated, err := updateQ.Save(ctx.Context())
	if err != nil {
		if ent.IsConstraintError(err) {
			return ctx.Status(400).JSON("Product with this name already exists")
		}

		return ctx.Status(500).JSON(err.Error())
	}

	return ctx.Status(200).JSON(updated)
}
