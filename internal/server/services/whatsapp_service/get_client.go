package whatsapp_service

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func (s *service) GetClient(phone string) (*whatsmeow.Client, error) {
	client := s.connectedClients[phone]

	if client != nil {
		return client, nil
	}

	userJid := types.NewJID(phone, types.DefaultUserServer)
	deviceStore, err := s.store.GetDevice(userJid)
	if deviceStore == nil {
		deviceStore = s.store.NewDevice()
	}

	if err != nil {
		return nil, err
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)

	client.AddEventHandler(s.EventHandler)

	return client, nil
}
