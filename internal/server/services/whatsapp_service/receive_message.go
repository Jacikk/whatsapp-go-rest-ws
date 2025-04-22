package whatsapp_service

type Message struct {
	From        string
	Receiver    string
	MessageType string
	Message     *string
	IsFromMe    bool
	IsGroup     bool
}

func (s *service) SendToMessagesChannel(message Message) error {
	if s.MessagesChan == nil {
		s.MessagesChan = make(chan Message)
	}

	s.MessagesChan <- message
	return nil
}

func (s *service) GetMessagesChannel() chan Message {
	return s.MessagesChan
}
