package whatsapp_service

import (
	"whatsapp_rest_ws/internal/server/errs"

	"go.mau.fi/whatsmeow/types"
)

func (s *service) GetGroups(phone string) ([]*types.GroupInfo, error) {
	client, err := s.GetClient(phone)
	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, errs.ClientNotFound(phone)
	}

	if !client.IsConnected() {
		return nil, errs.ClientNotConnected(phone)
	}

	if !client.IsLoggedIn() {
		return nil, errs.ClientNotLoggedIn(phone)
	}

	groups, err := client.GetJoinedGroups()
	if err != nil {
		return nil, err
	}

	return groups, nil
}
