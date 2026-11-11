package main

import (
	"fmt"
	"github.com/montacore/blog_aggregator/internal/config"
)

type state struct {
	cfg			Config

}
func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading from config file: %s", err)
	}
	err = cfg.SetUser("montacore")
	if err != nil {
		fmt.Printf("error setting username: %s", err)
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("error reading new config data: %s", err)
	}
	fmt.Println(cfg)
}
