package whatsapp_service

import (
	"fmt"

	"go.mau.fi/whatsmeow"
)

func (s *service) GetPairingCode(phone string, client *whatsmeow.Client) (*string, error) {
	if client == nil {
		return nil, fmt.Errorf("client not found for phone: %s", phone)
	}

	pairingCode, err := client.PairPhone(phone, true, whatsmeow.PairClientChrome, "Chrome (MacOS)")
	if err != nil {
		return nil, err
	}

	return &pairingCode, nil
}
