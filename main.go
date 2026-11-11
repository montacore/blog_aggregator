package main

import (
	"fmt"
	"github.com/montacore/blog_aggregator/internal/config"
	"os"
)

type state struct {
	cfg			*config.Config

}
func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading from config file: %s", err)
		os.Exit(1)
	}
	programState := state{&cfg}

	if len(os.Args) <= 2 {
		fmt.Println("[ERROR] not enough arguments provided")
		os.Exit(1)
	}
	cmd := command{
		command: os.Args[1],
		args: os.Args[2:],
	}

	cmds := commands{ make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	err  = cmds.run(&programState, cmd)
	if err != nil {
		fmt.Printf("[error] encountered error running handler: %s\n", err)
	}

}
