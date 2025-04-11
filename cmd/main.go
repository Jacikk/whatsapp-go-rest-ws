package main

import (
	"bd_test/internals/env"
	"bd_test/internals/server"
)

func main() {
	if err := env.ValidateEnv(); err != nil {
		panic("error validating env: " + err.Error())
	}
	port := env.PORT
	server.New(port)
}
