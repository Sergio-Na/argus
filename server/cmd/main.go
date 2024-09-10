package main

import (
	"log"

	"github.com/Sergio-Na/argus/server/config"
	"github.com/Sergio-Na/argus/server/internal/server"
	"github.com/Sergio-Na/argus/server/internal/supabase"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	supaClient, err := supabase.NewSupabaseClient(cfg.SupabaseURL, cfg.SupabaseKey)
	if err != nil {
		log.Fatalf("Failed to create Supabase client: %v", err)
	}

	s, err := server.New(cfg, supaClient)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	if err := s.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
