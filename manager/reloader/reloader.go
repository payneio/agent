// reloader will watch the agent process
// If the process dies, it will restart it.
package reloader

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Run(pid int, command string, args []string) {

	change := make(chan bool)

	go watch(pid, change)
	for {
		<-change
		reload(pid, command, args)
	}

}

// RunWIthCmd is a convenience method for Run() which allows passing in
// the command and args as a single command string.
func RunWithCmd(pid int, cmd string) {

	// TODO: Handle escaped quotes (nested args)
	command_array := strings.Split(cmd, " ")
	command := command_array[0]
	args := command_array[1:len(command_array)]

	Run(pid, command, args)

}

// watch continuously polls every second for the
// existence of the given process.
func watch(pid int, change chan bool) {

	for {

		time.Sleep(1 * time.Second)

		//  Poll existience of new file
		_, err := os.FindProcess(pid)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		// change <- true
		// break

	}

}

// reload handles the dirty work of stopping and restarting a process.
// It will stop the given pid, swap files and issue the given restart command.
func reload(pid int, command string, args []string) {

	log.Println("Agent died! Restarting.")

	// restart
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)

	// TODO: get new pid

	Run(pid, command, args)

}
