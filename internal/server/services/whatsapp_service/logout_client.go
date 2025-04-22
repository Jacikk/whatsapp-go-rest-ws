package whatsapp_service

func (s *service) LogoutClient(phone string) error {
	client := s.connectedClients[phone]

	if client != nil {
		err := client.Logout()
		if err != nil {
			return err
		}
		delete(s.connectedClients, phone)
	}

	return nil
}
