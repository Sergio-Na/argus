package supabase

import (
	supa "github.com/supabase-community/supabase-go"
)

type SupabaseClient struct {
	Client *supa.Client
}

func NewSupabaseClient(supabaseURL, supabaseKey string) (*SupabaseClient, error) {
	client, err := supa.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		return nil, err
	}
	return &SupabaseClient{Client: client}, nil
}
