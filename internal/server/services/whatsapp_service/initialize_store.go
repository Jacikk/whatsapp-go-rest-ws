package whatsapp_service

import (
	"fmt"
	"os"
	"os/signal"
	"whatsapp_rest_ws/internal/env"

	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"

	_ "github.com/lib/pq"
)

func (s *service) InitializeStore() (*sqlstore.Container, error) {
	if s.store != nil {
		return nil, fmt.Errorf("client store is already initialized")
	}

	dbLog := waLog.Stdout("Database", "DEBUG", true)

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		env.PG_USER, env.PG_PASSWORD, env.PG_HOST, env.PG_PORT, env.PG_DBNAME, env.PG_SSLMODE)

	container, err := sqlstore.New("postgres", connString, dbLog)
	if err != nil {
		return nil, err
	}

	// Add listener for Ctrl+C to disconnect the store
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		fmt.Println("Disconnecting store...")
		container.Close()
		os.Exit(0)
	}()

	return container, nil
}
