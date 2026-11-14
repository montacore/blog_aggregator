package main

import (
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"
	"github.com/montacore/blog_aggregator/internal/database"
)
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("[error] not enough arguments provided")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("[error] encountered retrieving username %w", err)
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("[error] error setting user: %w", err)
	}
	fmt.Println("User successfully set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("must include name parameter")
	}
	now := time.Now()
	name := cmd.args[0]
	params:= database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name: name,
	}

	newRecord, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(newRecord.Name)
	if err != nil {
		return err
	}

	fmt.Println("created successfully:", newRecord)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("failed to clear users: %w", err)
	}
	fmt.Println("Users successfully cleared!")
	return nil
}

func handlerUsers(s *state, cmd command) error{
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("[error] retrieving users data from table, %w", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUsername {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("[error] failed to fetch RSS feed data: %w", err)
	}
	fmt.Printf("%+v",feed)
	return nil
}
