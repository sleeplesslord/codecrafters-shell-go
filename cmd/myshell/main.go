package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func exitCommand(args string) {
	exitCode, _ := strconv.Atoi(args)
	os.Exit(exitCode)
}

func echoCommand(args string) {
	fmt.Println(args)
}

func typeCommand(args string) {
	if _, ok := BuiltInHandlers[args]; ok {
		fmt.Printf("%s is a shell builtin\n", args)
	} else {
		fmt.Printf("%s not found\n", args)
	}
}

var BuiltInHandlers map[string]func(string)

func handleCommand(command string, args string) {
	if handler, ok := BuiltInHandlers[command]; ok {
		handler(args)
		return
	}

	fmt.Printf("%s: command not found\n", command)
}

func init() {
	BuiltInHandlers = map[string]func(string){
		"exit": exitCommand,
		"echo": echoCommand,
		"type": typeCommand,
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		command, args, _ := strings.Cut(input, " ")

		handleCommand(strings.Trim(command, "\n"), strings.Trim(args, "\n"))
	}
}
