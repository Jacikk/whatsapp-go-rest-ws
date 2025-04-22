package errs

import (
	"fmt"
	"strings"
)

const (
	clientNotFoundError         = "client not found"
	whatsappClientNotFoundError = "whatsapp client not found"
	clientNotConnectedError     = "client not connected"
	clientNotLoggedInError      = "client not logged in"
)

func ClientNotFound(phone string) error {
	return fmt.Errorf("%s for phone %s", clientNotFoundError, phone)
}

func IsClientNotFoundError(err error) bool {
	return strings.Contains(err.Error(), clientNotFoundError)
}

func WhatsappClientNotFound(phone string) error {
	return fmt.Errorf("%s for phone %s", whatsappClientNotFoundError, phone)
}

func IsWhatsappClientNotFoundError(err error) bool {
	return strings.Contains(err.Error(), whatsappClientNotFoundError)
}

func ClientNotConnected(phone string) error {
	return fmt.Errorf("%s for phone %s", clientNotConnectedError, phone)
}

func IsClientNotConnectedError(err error) bool {
	return strings.Contains(err.Error(), clientNotConnectedError)
}

func ClientNotLoggedIn(phone string) error {
	return fmt.Errorf("%s for phone %s", clientNotLoggedInError, phone)
}

func IsClientNotLoggedInError(err error) bool {
	return strings.Contains(err.Error(), clientNotLoggedInError)
}

func InvalidParameter(parameter string) error {
	return fmt.Errorf("invalid parameter: %s", parameter)
}

func IsInvalidParameterError(err error) bool {
	return strings.Contains(err.Error(), "invalid parameter")
}
