package main

import (
	"github.com/payneio/agent/config"
	"log"
)

type Agent struct {
	config config.CfgAgent
}

var agent Agent

func main() {

	config.ConfigFilename = "../conf/agent.toml"
	err := config.Load()
	if err != nil {
		log.Fatalf("I could not load my configuration file: %v", err)
	}
	agent.config = config.Config
	log.Println(agent.config.Process.Start)

	log.Println("Hello World.")

}
