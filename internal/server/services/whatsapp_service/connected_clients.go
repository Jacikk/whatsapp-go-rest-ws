package whatsapp_service

import "go.mau.fi/whatsmeow"

func (s *service) GetConnectedClients() map[string]*whatsmeow.Client {
	return s.connectedClients
}
