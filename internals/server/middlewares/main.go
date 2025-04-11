package middlewares

import (
	"bd_test/internals/server/ent"

	"github.com/gofiber/fiber/v3"
)

type basicAuth struct {
	Db *ent.Client
}

type BasicAuth interface {
	AccessKeyMiddleware() fiber.Handler
}

func New(db *ent.Client) BasicAuth {
	return &basicAuth{
		Db: db,
	}
}
