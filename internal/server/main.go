package server

import (
	"fmt"
	"log"
	"net/http"
	"whatsapp_rest_ws/internal/server/controllers/clients_controller"
	"whatsapp_rest_ws/internal/server/routes"

	"github.com/gofiber/fiber/v3"
	socketio "github.com/googollee/go-socket.io"
	_ "github.com/lib/pq"
)

func New(port int) {
	app := fiber.New()

	// Create the Socket.IO server
	ioServer := socketio.NewServer(nil)

	ioServer.OnConnect("/", func(s socketio.Conn) error {
		log.Println("[Socket.IO] Connected:", s.ID())
		return nil
	})

	ioServer.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		log.Println("[Socket.IO] Received:", msg)
		s.Emit("reply", "got your message: "+msg)
	})

	ioServer.OnError("/", func(s socketio.Conn, err error) {
		log.Println("[Socket.IO] Error:", err)
	})

	ioServer.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("[Socket.IO] Disconnected:", reason)
	})

	// Socket.IO server (WebSocket)
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/socket.io/", ioServer)

		addr := fmt.Sprintf(":%d", port+1) // Example: 3001 if Fiber is 3000
		log.Printf("[Socket.IO] Listening on %s\n", addr)
		log.Fatal(http.ListenAndServe(addr, mux))
	}()

	// Set up your normal Fiber routes
	clientsCtrler := clients_controller.New(ioServer)
	routes.RegisterClientsRouter(app, clientsCtrler)

	// Fiber server (HTTP API)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
