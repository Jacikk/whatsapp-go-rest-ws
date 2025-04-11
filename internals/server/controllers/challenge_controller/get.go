package challenge_controller

import "github.com/gofiber/fiber/v3"

func (c *controllers) Get(ctx fiber.Ctx) error {
	data := map[string]interface{}{
		"Teste Prático de Front-end": map[string]interface{}{
			"Objetivo": []string{
				"UI/UX",
				"Organização e estrutura do código",
				"Uso de Git (preferencialmente via GitHub)",
				"Capacidade de resolver problemas",
				"Clareza no feedback e sugestões de melhoria",
			},
		},
		"Desafio": map[string]interface{}{
			"Descrição": "Desenvolver uma aplicação front-end com funcionalidades de CRUD para produtos e categorias.",
			"Produtos": []string{
				"Listagem com filtros por nome e categoria",
				"Criação de novo produto",
				"Edição do nome do produto",
				"Exclusão de produto",
				"Página de detalhes do produto",
			},
			"Categorias": []string{
				"CRUD básico (listagem, criação, edição e exclusão)",
			},
			"Observação": "A aplicação deve ser capaz de realizar o build e estar pronta para implantação.",
		},
		"Requisitos obrigatórios": []string{
			"React com TypeScript",
			"CRUD completo de produtos",
			"CRUD completo de categorias",
		},
		"Diferenciais": map[string]interface{}{
			"Tecnologias": []string{"Next.js"},
			"Documentação": []string{
				"Instruções de build e implantação",
				"Prototipagem de baixa fidelidade",
				"Prototipagem de alta fidelidade (preferencialmente via Figma)",
			},
		},
		"Dicas": []string{
			"Utilize variáveis de ambiente para URLs da API e chaves de acesso",
			"Faça commits atômicos, pequenos e objetivos a cada nova funcionalidade ou ajuste",
			"Use commits semânticos: feat, fix, chore, style, docs, etc., com estados opcionais como wip, done, rollback",
			"Use bibliotecas de componentes como Material UI ou Ant Design",
		},
	}

	return ctx.Status(200).JSON(data)
}
