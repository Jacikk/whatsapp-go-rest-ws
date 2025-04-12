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

var instance *basicAuth

func New(db *ent.Client) BasicAuth {
	if instance != nil {
		return instance
	}
	
	instance = &basicAuth{
		Db: db,
	}
	return instance
}
