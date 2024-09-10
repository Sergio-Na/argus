package database

import (
	"github.com/Sergio-Na/argus/server/internal/supabase"
	"github.com/supabase-community/gotrue-go/types"
)

type Service struct {
	client *supabase.SupabaseClient
}

func NewService(client *supabase.SupabaseClient) *Service {
	return &Service{client: client}
}

func (s *Service) GetUser() (*types.UserResponse, error) {
	user, err := s.client.Client.Auth.GetUser()
	return user, err
}
