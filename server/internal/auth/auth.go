package auth

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

func (s *Service) SignUp(email, password string) error {
	request := types.SignupRequest{
		Email:    email,
		Password: password,
	}
	_, err := s.client.Client.Auth.Signup(request)
	return err
}

func (s *Service) SignIn(email string, password string) (types.TokenResponse, error) {
	TokenResponse, err := s.client.Client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		return types.TokenResponse{}, err
	}
	return *TokenResponse, nil
}

func (s *Service) VerifyToken(token string) (*types.UserResponse, error) {
	user, err := s.client.Client.Auth.WithToken(token).GetUser()
	return user, err
}
