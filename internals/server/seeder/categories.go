package seeder

import (
	"bd_test/internals/server/ent"
	"context"
	"log"
)

func SeedCategories(client *ent.Client) {
	ctx := context.Background()

	defaultCategories := map[int]string{
		1: "Eletrônicos",
		2: "Cosméticos",
		3: "Perfumes",
	}

	for id, name := range defaultCategories {
		c, err := client.Category.Get(ctx, id)

		if ent.IsNotFound(err) {
			// Category does not exist, create it
			_, err := client.Category.
				Create().
				SetName(name).
				Save(ctx)
			if err != nil {
				log.Printf("❌ Failed to create category %d - %q: %v", id, name, err)
			} else {
				log.Printf("✅ Created category %d: %s", id, name)
			}
		} else if err != nil {
			log.Printf("❌ Error checking category %d: %v", id, err)
		} else if c.Name != name {
			// Category exists but has wrong name, update it
			_, err := client.Category.
				UpdateOneID(id).
				SetName(name).
				Save(ctx)
			if err != nil {
				log.Printf("❌ Failed to update category %d to %q: %v", id, name, err)
			} else {
				log.Printf("🛠️  Updated category %d to name: %s", id, name)
			}
		} else {
			log.Printf("✅ Category %d already exists with correct name: %s", id, name)
		}
	}
}
