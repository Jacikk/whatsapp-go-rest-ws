package whatsapp_service

import (
	"whatsapp_rest_ws/internal/server/errs"

	"go.mau.fi/whatsmeow/types"
)

func (s *service) GetContacts(phone string) (map[types.JID]types.ContactInfo, error) {
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

	contacts, err := client.Store.Contacts.GetAllContacts()
	if err != nil {
		return nil, err
	}

	return contacts, nil
}
