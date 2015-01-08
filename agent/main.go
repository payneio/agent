package main

import (
	"github.com/payneio/agent/config"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
)

var agent = Agent{}

type Agent struct {
	config config.Configuration
}

func main() {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals
		cleanup()
		os.Exit(1)
	}()

	select {}

}

func init() {
	config.ConfigFilename = "../conf/agent.toml"
	if err := config.Load(); err != nil {
		log.Fatalf("I could not load my configuration file: %v", err)
	}
	agent.config = config.Config
	log.Printf("Loaded configuration v.%s\n", agent.config.Version)

	if err := writePid(); err != nil {
		log.Fatalf("I could not write my PID file: %v", err)
	}
	log.Printf("Saved PID file to %s\n", agent.config.Process.PidFile)

}

func writePid() error {
	pid := []byte(strconv.Itoa(os.Getpid()))
	if err := ioutil.WriteFile(agent.config.Process.PidFile, pid, 0644); err != nil {
		return err
	}
	return nil
}

func cleanup() {
	if err := os.Remove(agent.config.Process.PidFile); err != nil {
		log.Printf("Could not delete PID file: %v, %v\n", agent.config.Process.PidFile, err)
	} else {
		log.Println("Deleted PID file.")
	}
}
