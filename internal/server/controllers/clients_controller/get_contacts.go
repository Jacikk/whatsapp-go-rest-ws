package clients_controller

import (
	"whatsapp_rest_ws/internal/server/errs"

	"github.com/gofiber/fiber/v3"
)

type ContactDto struct {
	Jid  string `json:"jid"`
	Name string `json:"name"`
}

type GetContactsResponse struct {
	Contacts []*ContactDto `json:"contacts"`
}

func (c *controller) GetContacts(ctx fiber.Ctx) error {
	phone := ctx.Params("phone")
	if phone == "" {
		return errs.InvalidParameter("phone is required")
	}

	client, err := c.whatsappService.GetClient(phone)
	if err != nil {
		return err
	}

	if client == nil {
		return errs.ClientNotFound(phone)
	}

	contacts, err := c.whatsappService.GetContacts(phone)
	if err != nil {
		return err
	}

	var cts = make([]*ContactDto, 0, len(contacts))

	for contactJid, contact := range contacts {
		cts = append(cts, &ContactDto{
			Jid:  contactJid.String(),
			Name: contact.FullName,
		})
	}

	return ctx.Status(200).JSON(cts)
}
