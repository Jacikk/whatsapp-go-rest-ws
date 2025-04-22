package whatsapp_service

import (
	"fmt"
	"strings"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type Service interface {
	GetClient(phone string) (*whatsmeow.Client, error)
	GetPairingCode(phone string, client *whatsmeow.Client) (*string, error)
	GetContacts(phone string) (map[types.JID]types.ContactInfo, error)
	GetGroups(phone string) ([]*types.GroupInfo, error)
	SendMessage(from string, to string, message string) error
	GetConnectedClients() map[string]*whatsmeow.Client
	LogoutClient(phone string) error

	SendToStatusChannel(status Status) error
	GetStatusChannel() chan Status

	SendToMessagesChannel(message Message) error
	GetMessagesChannel() chan Message
}

type service struct {
	store            *sqlstore.Container
	connectedClients map[string]*whatsmeow.Client

	StatusChan   chan Status
	MessagesChan chan Message
}

var whatsappService *Service

func New() Service {
	if whatsappService != nil {
		return *whatsappService
	}

	s := service{
		connectedClients: make(map[string]*whatsmeow.Client),
		StatusChan:       make(chan Status),
		MessagesChan:     make(chan Message),
	}

	store, err := s.InitializeStore()
	if err != nil {
		panic(err)
	}
	s.store = store

	devices, err := store.GetAllDevices()
	if err != nil {
		panic(err)
	}

	for _, device := range devices {
		client := whatsmeow.NewClient(device, waLog.Stdout("Client", "DEBUG", true))
		client.AutomaticMessageRerequestFromPhone = true

		client.AddEventHandler(s.EventHandler)
		x := strings.Split(client.Store.ID.String(), ":")

		s.connectedClients[x[0]] = client

		if client.Store.ID != nil {
			client.Connect()
		}
	}

	return &s
}

func (s *service) EventHandler(evt any) {
	switch v := evt.(type) {
	case *events.Message:
		var msg *string

		if v.Message.ExtendedTextMessage.Text != nil {
			msg = v.Message.ExtendedTextMessage.Text
		}

		if v.Message.Conversation != nil {
			msg = v.Message.Conversation
		}

		s.SendToMessagesChannel(Message{
			From:        v.Info.Sender.User,
			Receiver:    v.Info.Chat.User,
			Message:     msg,
			IsFromMe:    v.Info.IsFromMe,
			IsGroup:     v.Info.IsGroup,
			MessageType: v.Info.Type,
		})
	case *events.UndecryptableMessage:
		fmt.Println("Undecryptable message:", v)

	default:
		fmt.Printf("Received an event of type: %T\n", v)
	}

}
