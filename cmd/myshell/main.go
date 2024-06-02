package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func handleCommand(command string, args string) (ok bool) {
	ok = handleBuiltIn(command, args)
	if ok {
		return
	}

	return handleExternalCommand(command, args)
}

func handleExternalCommand(command string, args string) (ok bool) {
	var commandPath string
	var runnable bool

	if filepath.IsAbs(command) {
		commandPath = command
		runnable = true
	}

	if path, found := resolveFromPathVariable(command); found {
		commandPath = filepath.Join(path, command)
		runnable = true
	}

	cmd := exec.Command(commandPath, strings.Fields(args)...)
	output, _ := cmd.CombinedOutput()
	fmt.Printf("%s", output)
	return runnable
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		command, args, _ := strings.Cut(input, " ")
		command = strings.Trim(command, "\n")
		args = strings.Trim(args, "\n")

		handled := handleCommand(command, args)
		if !handled {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
