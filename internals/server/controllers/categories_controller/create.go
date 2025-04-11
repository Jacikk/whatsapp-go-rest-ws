package categories_controller

import (
	"bd_test/internals/server/ent"

	"github.com/gofiber/fiber/v3"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func (c controllers) Create(ctx fiber.Ctx) error {
	req := new(CreateCategoryRequest)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	accessKeyID := ctx.Locals("access_key_id").(int)

	if accessKeyID == 0 {
		ctx.Status(400).SendString("Access Key ID must be a valid integer")
		return nil
	}

	if req.Name == "" {
		ctx.Status(400).SendString("Name is required")
		return nil
	}

	tx, err := c.Db.Tx(ctx.Context())
	if err != nil {
		ctx.Status(500).SendString(err.Error())
		return nil
	}

	defer tx.Rollback()

	cat, err := tx.Category.Create().
		SetName(req.Name).
		SetAccessKeyID(accessKeyID).
		Save(ctx.Context())

	if err != nil {
		if ent.IsConstraintError(err) {
			ctx.Status(400).SendString("Category with this name already exists")
			return nil
		}

		ctx.Status(500).SendString(err.Error())
		return nil
	}

	if err := tx.Commit(); err != nil {
		ctx.Status(500).SendString(err.Error())
		return nil
	}

	return ctx.Status(200).JSON(cat)
}
