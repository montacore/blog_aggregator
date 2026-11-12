package main

import (
	"github.com/montacore/blog_aggregator/internal/config"
	"os"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/montacore/blog_aggregator/internal/database"
	"log"
)

type state struct {
	cfg			*config.Config
	db			*database.Queries

}
func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("[error] there was a problem reading the config file: %w", err)
	}
	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	dbNew := database.New(db)
	if err != nil {
		log.Fatalf("[ERROR] there was an issue creating database connection %s", err)

	}
	programState := state{&cfg, dbNew}
	if len(os.Args) < 2 {
		log.Fatal("you must provide, at minimum, a command: %w", err)
	}
	cmd := command{
		command: os.Args[1],
		args: os.Args[2:],
	}

	cmds := commands{ make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	err  = cmds.run(&programState, cmd)
	if err != nil {
		log.Fatal("[error] encountered error registering command: %w", err)
	}
	

}
