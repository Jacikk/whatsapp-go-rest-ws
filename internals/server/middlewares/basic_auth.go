package middlewares

import (
	"bd_test/internals/server/ent"
	"bd_test/internals/server/ent/accesskeys"

	"github.com/gofiber/fiber/v3"
)

func (b *basicAuth) AccessKeyMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		accessKey := c.Get("Access-Key")
		if accessKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON("Access key is required")
		}

		ak, err := b.Db.AccessKeys.Query().Where(accesskeys.KeyContainsFold(accessKey)).Only(c.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusUnauthorized).JSON("Invalid access key")
			}
			return c.Status(fiber.StatusInternalServerError).JSON("Error checking access key")
		}

		c.Locals("access_key_id", ak.ID)

		// Save the path and access_key_id in Logs
		b.Db.Logs.Create().
			SetNillableRoute(&c.Route().Path).
			SetNillableMethod(&c.Route().Method).
			SetAccessKeyID(ak.ID).
			Save(c.Context())

		return c.Next()
	}
}
