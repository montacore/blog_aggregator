package main

import (
	"fmt"
)
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("[error] no argument provided")
	}

}
