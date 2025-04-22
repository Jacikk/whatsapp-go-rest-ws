package whatsapp_service

type Status struct {
	PhoneNumber string `json:"phone"`
	Status      string `json:"status"`
}

const (
	Authenticated  = "Authenticated"
	Connected      = "Connected"
	Disconnected   = "Disconnected"
	LoggedOut      = "Logged out"
	ConnectFailure = "Connection failed"
	HasClient      = "Has client"
)

func (s *service) SendToStatusChannel(status Status) error {
	if s.StatusChan == nil {
		s.StatusChan = make(chan Status)
	}

	s.StatusChan <- status
	return nil
}

func (s *service) GetStatusChannel() chan Status {
	return s.StatusChan
}
