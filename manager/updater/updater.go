// updater will watch for a new version of a binary (filename.new).
// When it sees a new version, it will stop the process, swap the binaries,
// and restart a new process.
package updater

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Run takes a pid of a currently running process,
// a filename to watch for a new version (filename.new), and
// the command to run once it sees it (usually a start) along with the
// command's args.
func Run(pid int, filename string, command string, args []string) {

	change := make(chan bool)

	go watch(filename, change)
	for {
		<-change
		reload(pid, filename, command, args)
	}

}

// RunWIthCmd is a convenience method for Run() which allows passing in
// the command and args as a single command string.
func RunWithCmd(pid int, filename string, cmd string) {

	// TODO: Handle escaped quotes (nested args)
	command_array := strings.Split(cmd, " ")
	command := command_array[0]
	args := command_array[1:len(command_array)]

	Run(pid, filename, command, args)

}

// watch continuously polls every second for the
// existence of the given filename.
func watch(filename string, change chan bool) {

	newFile := fmt.Sprint(filename, ".new")

	for {

		//  Poll existience of new file
		_, err := os.Stat(newFile)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		change <- true
		break

	}

}

// reload handles the dirty work of stopping and restarting a process.
// It will stop the given pid, swap files and issue the given restart command.
func reload(pid int, filename string, command string, args []string) {

	log.Println("File changed. Restarting.")

	// kill existing pid
	// TODO

	// rename existing file
	// TODO

	// rename new file
	// TODO

	// restart
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)

	// TODO: get new pid
	Run(pid, filename, command, args)

}
