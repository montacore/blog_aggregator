package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	"github.com/montacore/blog_aggregator/internal/database"
	"context"
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

