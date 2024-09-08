package main

import (
	"log"

	"github.com/Sergio-Na/argus/server/config"
	"github.com/Sergio-Na/argus/server/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	s, err := server.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	if err := s.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
