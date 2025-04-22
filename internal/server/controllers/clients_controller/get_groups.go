package clients_controller

import (
	"whatsapp_rest_ws/internal/server/errs"

	"github.com/gofiber/fiber/v3"
)

type GroupDto struct {
	Jid  string `json:"jid"`
	Name string `json:"name"`
}

type GetGroupsResponse struct {
	Groups []*GroupDto `json:"groups"`
}

func (c *controller) GetGroups(ctx fiber.Ctx) error {
	phone := ctx.Params("phone")
	client, err := c.whatsappService.GetClient(phone)
	if err != nil {
		return err
	}

	if client == nil {
		return errs.ClientNotFound(phone)
	}

	groups, err := c.whatsappService.GetGroups(phone)
	if err != nil {
		return err
	}

	var grps = make([]*GroupDto, 0, len(groups))

	for _, group := range groups {
		grps = append(grps, &GroupDto{
			Jid:  group.JID.String(),
			Name: group.Name,
		})
	}

	return ctx.Status(200).JSON(grps)
}
