package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
	commandPath, runnable := findExecutableFile(command)

	if runnable {
		cmd := exec.Command(commandPath, strings.Fields(args)...)
		output, _ := cmd.CombinedOutput()
		fmt.Printf("%s", output)
	}

	return runnable
}

func main() {
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
