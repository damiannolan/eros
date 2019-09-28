package main

import (
	"github.com/damiannolan/eros/health"
	log "github.com/damiannolan/eros/logger"
	"github.com/damiannolan/eros/server"
)

func main() {
	srv := server.New(server.NewConfig())
	srv.RegisterResource(health.NewResource("/health"))

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("Serving failed")
	}
}
