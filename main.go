package main

import (
	"fmt"
	"github.com/montacore/blog_aggregator/internal/config"
	"os"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/montacore/blog_aggregator/internal/database"
)

type state struct {
	cfg			*config.Config
	db			*database.Queries

}
func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading from config file: %s", err)
		os.Exit(1)
	}
	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	dbNew := database.New(db)
	if err != nil {
		fmt.Printf("[ERROR] there was an issue creating database connection %s", err)
	}
	programState := state{&cfg, dbNew}
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
	cmds.register("register", handlerRegister)
	err  = cmds.run(&programState, cmd)
	if err != nil {
		fmt.Printf("[error] encountered error running handler: %s\n", err)
		os.Exit(1)
	}
	

}
