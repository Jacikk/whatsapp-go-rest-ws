package products_controller

import (
	"bd_test/internals/server/ent"

	"github.com/gofiber/fiber/v3"
)

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Price       float64 `json:"price"`
	CategoryId  int     `json:"category_id"`
}

func (c controllers) Create(ctx fiber.Ctx) error {
	req := CreateProductRequest{
		Name:        fiber.Query(ctx, "name", ""),
		Description: fiber.Query(ctx, "description", ""),
		ImageUrl:    fiber.Query(ctx, "image_url", ""),
		Price:       fiber.Query(ctx, "price", 0.0),
		CategoryId:  fiber.Query(ctx, "category_id", 0),
	}

	accessKeyID := ctx.Locals("access_key_id").(int)
	if accessKeyID == 0 {
		ctx.Status(400).SendString("Access Key ID is required")
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

	cat, err := tx.Product.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		SetImageURL(req.ImageUrl).
		SetPrice(req.Price).
		SetCategoryID(req.CategoryId).
		SetAccessKeyID(accessKeyID).
		Save(ctx.Context())

	if err != nil {
		if ent.IsConstraintError(err) {
			ctx.Status(400).SendString("Product with this name already exists")
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
