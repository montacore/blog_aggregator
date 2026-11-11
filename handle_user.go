package main

import (
	"fmt"
)
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("[error] not enough arguments provided")
	}
	 err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("[error] error setting user: %w", err)
	}
	fmt.Println("User successfully set")
	return nil
}
