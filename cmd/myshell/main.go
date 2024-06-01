package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var pathDirectories []string

func handleCommand(command string, args string) (ok bool) {
	ok = handleBuiltIn(command, args)
	if ok {
		return
	}

	return handleExternalCommand(command, args)
}

func handleExternalCommand(command string, args string) (ok bool) {
	if path, found := fileInPathVariables(command); found {
		execPath := fmt.Sprintf("%s/%s", path, command)
		cmd := exec.Command(execPath, args)
		output, _ := cmd.Output()
		fmt.Printf("%s", output)
		return true
	}

	return false
}

func fileInPath(path string, command string) (found bool) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return false
	}

	for _, e := range entries {
		if e.Name() == command {
			return true
		}
	}

	return false
}

func fileInPathVariables(command string) (path string, found bool) {
	for _, path := range pathDirectories {
		if fileInPath(path, command) {
			return path, true
		}
	}

	return "", false
}

func init() {
	pathDirectories = strings.Split(os.Getenv("PATH"), ":")
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

		handled := handleCommand(command, strings.Trim(args, "\n"))
		if !handled {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
