package main

import (
	"github.com/payneio/agent/config"
	"log"
)

type Agent struct {
	config config.Configuration
}

var agent Agent

func main() {

	if err := init(); err != nil {
		log.Fatalf("I could not initialize: %v", err)
	}
	log.Printf("Loaded configuration v.%s\n", manager.config.Version)

}

func init() error {
	config.ConfigFilename = "../conf/agent.toml"
	if err := config.Load(); err != nil {
		return err
	}
	manager.config = config.Config
	return nil
}
