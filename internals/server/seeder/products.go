package seeder

import (
	"bd_test/internals/server/ent"
	"bd_test/internals/server/ent/product"
	"context"
	"log"
)

func SeedProducts(client *ent.Client) {
	ctx := context.Background()

	products := []struct {
		Name        string
		Price       float64
		Description string
		ImageURL    string
		CategoryID  int
	}{
		{"Smartphone X200", 1899.99, "Smartphone com câmera tripla e bateria de longa duração.", "https://images.unsplash.com/photo-1511707171634-5f897ff02aa9", 1},
		{"Fone de Ouvido Bluetooth", 299.90, "Fone sem fio com cancelamento de ruído ativo.", "https://images.unsplash.com/photo-1585386959984-a4155223f7c6", 1},
		{"Notebook Pro 15", 5799.00, "Notebook para profissionais com tela 15” e SSD de 512GB.", "https://images.unsplash.com/photo-1517336714731-489689fd1ca8", 1},
		{"Smartwatch FitBand", 499.99, "Relógio inteligente com monitoramento de saúde.", "https://images.unsplash.com/photo-1603791440384-56cd371ee9a7", 1},
		{"Câmera de Segurança Wi-Fi", 279.00, "Câmera com detecção de movimento e visão noturna.", "https://images.unsplash.com/photo-1591348277033-953d8f95b6d5", 1},
		{"Creme Facial Hidratante", 49.99, "Hidratação profunda para todos os tipos de pele.", "https://images.unsplash.com/photo-1588776814546-b2f6b68f7c7c", 2},
		{"Sabonete Natural de Lavanda", 19.90, "Sabonete artesanal com aroma suave de lavanda.", "https://images.unsplash.com/photo-1600180758890-6fbd2fbdfb6f", 2},
		{"Esfoliante Corporal Citrus", 34.90, "Renovação da pele com aroma cítrico refrescante.", "https://images.unsplash.com/photo-1600185365085-586b1c2b6a58", 2},
		{"Kit Cuidados com a Pele", 89.90, "Kit completo para limpeza e hidratação facial.", "https://images.unsplash.com/photo-1600181953606-1ec7b7bc9d97", 2},
		{"Shampoo Detox Herbal", 29.90, "Limpeza suave com ingredientes naturais.", "https://images.unsplash.com/photo-1598514982721-8d5f5b6a799f", 2},
		{"Perfume Aurora Feminino", 149.90, "Fragrância floral delicada e marcante.", "https://images.unsplash.com/photo-1617814081330-cf70fe8ef58b", 3},
		{"Perfume Noturno Masculino", 169.90, "Aroma amadeirado com notas de especiarias.", "https://images.unsplash.com/photo-1608231387042-66d1773070b4", 3},
		{"Colônia Fresh Day", 99.90, "Fragrância leve e refrescante para o dia a dia.", "https://images.unsplash.com/photo-1610814670498-2b6238fc9c45", 3},
		{"Perfume Lux Essence", 189.90, "Perfume de luxo com notas florais e cítricas.", "https://images.unsplash.com/photo-1605972781920-0903f6c15f37", 3},
		{"Body Splash Tropical", 59.90, "Splash corporal com notas de frutas tropicais.", "https://images.unsplash.com/photo-1621996346565-3c37758cc231", 3},
		{"Tablet 10” Multimídia", 1299.00, "Ideal para estudo, jogos e navegação.", "https://images.unsplash.com/photo-1587825140708-dfaf72ae4b04", 1},
		{"Roteador Dual Band", 149.00, "Internet rápida com tecnologia dual band.", "https://images.unsplash.com/photo-1580910051071-343b295d739b", 1},
		{"Máscara Facial de Argila", 24.90, "Limpeza profunda e controle de oleosidade.", "https://images.unsplash.com/photo-1588776814682-2944471d5f8c", 2},
		{"Condicionador Nutritivo", 27.90, "Hidratação e maciez para cabelos danificados.", "https://images.unsplash.com/photo-1600181953671-535d6aa30b20", 2},
		{"Perfume Oceano Azul", 129.90, "Aroma refrescante com notas aquáticas.", "https://images.unsplash.com/photo-1613459625606-4050d7ae3a4e", 3},
	}

	for _, p := range products {
		existing, err := client.Product.
			Query().
			Where(product.NameEQ(p.Name)).
			Only(ctx)

		if ent.IsNotFound(err) {
			// Create if not found
			_, err := client.Product.
				Create().
				SetName(p.Name).
				SetPrice(p.Price).
				SetDescription(p.Description).
				SetImageURL(p.ImageURL).
				SetCategoryID(p.CategoryID).
				Save(ctx)
			if err != nil {
				log.Printf("❌ Failed to create product %q: %v", p.Name, err)
			} else {
				log.Printf("✅ Created product: %s", p.Name)
			}
		} else if err != nil {
			log.Printf("❌ Error checking product %q: %v", p.Name, err)
		} else {
			// Check if any field needs to be updated
			needsUpdate := existing.Price != p.Price ||
				existing.Description != p.Description ||
				existing.ImageURL != p.ImageURL ||
				existing.CategoryID == nil || *existing.CategoryID != p.CategoryID

			if needsUpdate {
				update := client.Product.UpdateOneID(existing.ID).
					SetPrice(p.Price).
					SetDescription(p.Description).
					SetImageURL(p.ImageURL).
					SetCategoryID(p.CategoryID)

				_, err := update.Save(ctx)
				if err != nil {
					log.Printf("❌ Failed to update product %q: %v", p.Name, err)
				} else {
					log.Printf("🛠️  Updated product: %s", p.Name)
				}
			} else {
				log.Printf("✅ Product already up-to-date: %s", p.Name)
			}
		}
	}
}
