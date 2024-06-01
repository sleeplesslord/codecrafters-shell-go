package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		command, rest, _ := strings.Cut(input, " ")

		switch command {
		case "exit":
			exitCode, _ := strconv.Atoi(rest)
			os.Exit(exitCode)
		case "echo":
			fmt.Print(rest)
		default:
			fmt.Printf("%s: command not found\n", strings.Trim(command, "\n"))
		}
	}
}
