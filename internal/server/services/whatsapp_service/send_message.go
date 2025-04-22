package whatsapp_service

import (
	"context"
	"whatsapp_rest_ws/internal/server/errs"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
)

func (s *service) SendMessage(from string, to string, message string) error {
	client, err := s.GetClient(from)
	if err != nil {
		return err
	}

	if client == nil {
		return errs.ClientNotFound(from)
	}

	if !client.IsConnected() {
		return errs.ClientNotConnected(from)
	}

	targetJid := types.NewJID(to, types.DefaultUserServer)

	msg := waE2E.Message{
		Conversation: &message,
	}

	_, err = client.SendMessage(context.Background(), targetJid, &msg)
	if err != nil {
		return err
	}

	return nil
}
