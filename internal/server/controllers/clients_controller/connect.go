package clients_controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type ConnectRequest struct {
	Phone          string `json:"phone"`
	GetQrCode      *bool  `json:"get_qr_code"`
	GetPairingCode *bool  `json:"get_pairing_code"`
}

type ConnectResponse struct {
	QrCode      *string `json:"qr_code"`
	PairingCode *string `json:"pairing_code"`
	Status      string  `json:"status"`
	IsEnded     bool    `json:"is_ended"`
}

func (c *controller) Connect(ctx fiber.Ctx) error {
	req := new(ConnectRequest)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}

	req.Phone = ctx.Params("phone")
	if req.Phone == "" {
		return fiber.NewError(fiber.StatusBadRequest, "phone is required")
	}

	client, err := c.whatsappService.GetClient(req.Phone)
	if err != nil {
		return err
	}

	errChan := make(chan error, 1)

	if client.Store.ID == nil {
		qrCode, err := client.GetQRChannel(context.Background())
		if err != nil {
			return err
		}

		err = client.Connect()
		if err != nil {
			errChan <- err
		}
		if req.GetPairingCode != nil && *req.GetPairingCode {
			go func() {
				time.Sleep(time.Second * 2)
				pairingCode, err := c.whatsappService.GetPairingCode("+"+req.Phone, client)
				if err != nil {
					errChan <- err
					return
				}

				resp := ConnectResponse{
					PairingCode: pairingCode,
					Status:      "Awaiting for user to enter pairing code",
					IsEnded:     false,
				}

				if err := c.io.BroadcastToNamespace('/', ); err != nil {
					errChan <- err
					return
				}
			}()
		}

		go func() {
			for evt := range qrCode {
				if evt.Event == "code" && req.GetQrCode != nil && *req.GetQrCode {
					resp := ConnectResponse{
						QrCode:  &evt.Code,
						Status:  "Awaiting for user to scan QR code",
						IsEnded: false,
					}

					if err := server.Send(&resp); err != nil {
						errChan <- err
						return
					}
				}

				if evt.Event == "success" {
					resp := ConnectResponse{
						IsEnded: true,
						Status:  "Client connected successfully",
					}

					c.whatsappService.GetConnectedClients()[req.PhoneNumber] = client

					if err := server.Send(&resp); err != nil {
						errChan <- err
						return
					} else {
						errChan <- nil
					}
				}

				if evt.Event == "error" {
					errChan <- evt.Error
				}
			}
		}()
	} else {
		err = client.Connect()
		if err != nil {
			errChan <- err
		}

		resp := ConnectResponse{
			IsEnded: true,
			Status:  "Client connected",
		}

		if err := server.Send(&resp); err != nil {
			errChan <- err
		}

		errChan <- nil
	}

	if client.IsLoggedIn() {
		resp := ConnectResponse{
			IsEnded: true,
			Status:  "Client connected",
		}

		if err := server.Send(&resp); err != nil {
			errChan <- err
		}

		errChan <- nil
	}

	if err := <-errChan; err != nil {
		return err
	}

	return nil
}
