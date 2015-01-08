// agentManager is a process that runs alongside an agent
package main

import (
	"github.com/payneio/agent/config"
	"github.com/payneio/agent/manager/reloader"
	"github.com/payneio/agent/manager/updater"
	"log"
	"os"
)

var manager Manager

type Manager struct {
	config config.Configuration
}

func main() {

	manager = Manager{}

	runReloader()
	runUpdater()

	select {}

}

func init() {
	config.ConfigFilename = "../conf/agent.toml"
	if err := config.Load(); err != nil {
		log.Fatalf("I could not initialize: %v", err)
	}
	manager.config = config.Config
	log.Printf("Loaded configuration v.%s\n", manager.config.Version)
}

func runReloader() {
	pid := os.Getpid()
	// TODO: get PID of agent (from PID file, likely)
	go func() {
		reloader.RunWithCmd(pid, manager.config.Process.Start)
	}()
}

func runUpdater() {
	pid := os.Getpid()
	// TODO: get PID of agent (from PID file, likely)
	go func() {
		updater.RunWithCmd(pid, "agent", manager.config.Process.Stop)
	}()
}

// TODO: Handle SIGHUP: https://bitbucket.org/PinIdea/zero-downtime-daemon
