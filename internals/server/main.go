package server

import (
	"bd_test/internals/env"
	"bd_test/internals/server/controllers/categories_controller"
	"bd_test/internals/server/controllers/challenge_controller"
	"bd_test/internals/server/controllers/products_controller"
	"bd_test/internals/server/ent"
	"bd_test/internals/server/routes"
	"bd_test/internals/server/seeder"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
)

func New(port int) {
	app := fiber.New()

	connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		env.PG_HOST, env.PG_PORT, env.PG_USER, env.PG_DBNAME, env.PG_PASSWORD, env.PG_SSLMODE)

	client, err := ent.Open("postgres", connString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the seeder to populate the database with initial data.
	// This is useful for development and testing purposes.
	// In production, you might want to remove this or use a different approach for seeding data.
	// Must be in the same order as the migration.
	seeder.SeedCategories(client)
	seeder.SeedProducts(client)

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	challengeCtrler := challenge_controller.New()
	routes.RegisterChallengeRouter(app, challengeCtrler)

	categoriesCtrler := categories_controller.New(client)
	routes.RegisterCategoriesRouter(app, categoriesCtrler)

	productsCtrler := products_controller.New(client)
	routes.RegisterProductsRouter(app, productsCtrler)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
