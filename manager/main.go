// agentManager is a process that runs alongside an agent
package main

import (
	"flag"
	"github.com/payneio/agent/config"
	"github.com/payneio/agent/manager/reloader"
	"github.com/payneio/agent/manager/updater"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

var configFile = flag.String("config", "manager.toml", "The full path to the manager config file.")

var manager = Manager{}

type Manager struct {
	config config.Configuration
}

func main() {

	runReloader()
	runUpdater()

	select {}

}

func init() {
	config.ConfigFilename = *configFile
	if err := config.Load(); err != nil {
		log.Fatalf("I could not initialize: %v", err)
	}
	manager.config = config.Config
	log.Printf("Loaded configuration v.%s\n", manager.config.Version)
}

func runReloader() {

	go func() {
		for {
			pid, err := getPid()
			if err != nil {
				log.Printf("Could not get PID. Trying again. %v\n", err)
				time.Sleep(1 * time.Second)
			}
			reloader.RunWithCmd(pid, manager.config.Process.Start)
		}
	}()

}

func runUpdater() {
	pid := os.Getpid()
	// TODO: get PID of agent (from PID file, likely)
	go func() {
		updater.RunWithCmd(pid, "agent", manager.config.Process.Stop)
	}()
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

// TODO: Handle SIGHUP: https://bitbucket.org/PinIdea/zero-downtime-daemon
