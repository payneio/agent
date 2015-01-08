// agentManager is a process that runs alongside an agent
package main

import (
	"github.com/payneio/agent/config"
	"github.com/payneio/agent/manager/reloader"
	"github.com/payneio/agent/manager/updater"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var manager = Manager{}

type Manager struct {
	config config.Configuration
}

func main() {

	if err := runReloader(); err != nil {
		log.Fatalf("I could not start the reloader: %v", err)
	}
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

func getPid() (int, error) {
	log.Println(manager.config.Process.PidFile)
	pidBytes, err := ioutil.ReadFile(manager.config.Process.PidFile)
	if err != nil {
		return 0, err
	}
	pidS := string(pidBytes)
	pid, err := strconv.ParseUint(pidS, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(pid), nil
}

func runReloader() error {

	pid, err := getPid()
	if err != nil {
		return err
	}

	log.Printf("Got a PID for the reloader: %v\n", pid)
	go func() {
		reloader.RunWithCmd(pid, manager.config.Process.Start)
	}()
	return nil
}

func runUpdater() {
	pid := os.Getpid()
	// TODO: get PID of agent (from PID file, likely)
	go func() {
		updater.RunWithCmd(pid, "agent", manager.config.Process.Stop)
	}()
}

// TODO: Handle SIGHUP: https://bitbucket.org/PinIdea/zero-downtime-daemon
