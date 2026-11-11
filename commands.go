package main

type command struct {
	login			string
	args			[]string
}

type commands struct {
	command		 map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.commands[cmd.login]
	if !exists {
		return fmt.Errorf("no such command: %s", cmd.name )
	}
}
