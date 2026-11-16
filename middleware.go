package main

import (
	"context"
	"fmt"
	"github.com/montacore/blog_aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return fmt.Errorf("[error] error encountered retrieving user: %w", err)
		}
		
		return handler(s, cmd, user)
	}
}
