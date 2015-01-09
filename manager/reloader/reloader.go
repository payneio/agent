// reloader will watch the agent process
// If the process dies, it will restart it.
package reloader

import (
	"log"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func Run(pid int, command string, args []string) {

	change := make(chan bool, 1)

	if pid > 0 {
		go watch(pid, change)
		<-change
	}
	reload(command, args)

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

		// _, err := os.FindProcess(pid)

		// This works on POSIX only
		// https://groups.google.com/forum/#!topic/golang-nuts/hqrp0UHBK9k
		if err := syscall.Kill(pid, 0); err != nil {
			change <- true
			return
		}
		continue

	}

}

func reload(command string, args []string) {

	log.Println("Agent died! Restarting.")

	// restart the process
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)

}
