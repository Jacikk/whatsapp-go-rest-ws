package products_controller

import (
	"bd_test/internals/server/ent"
	"strings"

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
	req := new(CreateProductRequest)
	if err := ctx.Bind().Body(req); err != nil {
		return err
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

	if req.Description == "" {
		ctx.Status(400).SendString("Description is required")
		return nil
	}

	if req.ImageUrl == "" {
		ctx.Status(400).SendString("Image URL is required")
		return nil
	}

	if req.Price <= 0 {
		ctx.Status(400).SendString("Price must be greater than 0")
		return nil
	}

	if req.CategoryId <= 0 {
		ctx.Status(400).SendString("Category ID must be greater than 0")
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
			if strings.Contains(err.Error(), `violates foreign key constraint "products_categories_products"`) {
				ctx.Status(400).SendString("Category ID does not exist")
				return nil
			}

			if strings.Contains(err.Error(), `duplicate key value violates unique constraint "products_name_key"`) {
				ctx.Status(400).SendString("Product name already exists")
				return nil
			}
			ctx.Status(400).SendString(err.Error())
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
