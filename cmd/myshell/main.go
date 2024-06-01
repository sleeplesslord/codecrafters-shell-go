package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pathDirectories []string

func handleCommand(command string, args string) (ok bool) {
	return handleBuiltIn(command, args)
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

		handled := handleCommand(strings.Trim(command, "\n"), strings.Trim(args, "\n"))
		if !handled {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
