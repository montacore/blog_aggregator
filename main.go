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
		log.Fatalf("[error] there was a problem reading the config file: %v", err)
	}
	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	dbNew := database.New(db)
	if err != nil {
		log.Fatalf("[ERROR] there was an issue creating database connection %v", err)

	}
	programState := state{&cfg, dbNew}
	if len(os.Args) < 2 {
		log.Fatalf("you must provide, at minimum, a command: %v", err)
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
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	err  = cmds.run(&programState, cmd)
	if err != nil {
		log.Fatalf("[error] encountered error registering command: %v", err)
	}
	

}
