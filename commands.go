package main

import (
	"fmt"
)
type command struct {
	command			string
	args			[]string
}

type commands struct {
	commands		 map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.commands[cmd.command]
	if !exists {
		return fmt.Errorf("no such command: %s", cmd.command )
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
