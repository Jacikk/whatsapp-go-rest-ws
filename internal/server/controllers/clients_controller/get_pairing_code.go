package clients_controller

func (c *controller) GetPairingCode(ctx fiber.Ctx) error {
	clientID := ctx.Params("id")
	client, exists := c.whatsappService.GetConnectedClients()[clientID]
	if !exists {
		return ctx.Status(404).JSON(fiber.Map{"error": "Client not found"})
	}

	pairingCode, err := client.GetPairingCode()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to get pairing code"})
	}

	return ctx.JSON(fiber.Map{"pairing_code": pairingCode})
}