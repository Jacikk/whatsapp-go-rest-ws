package main

import (
	"whatsapp_rest_ws/internal/env"
	"whatsapp_rest_ws/internal/server"
)

func main() {
	if err := env.ValidateEnv(); err != nil {
		panic("error validating env: " + err.Error())
	}
	port := env.PORT
	server.New(port)
}
