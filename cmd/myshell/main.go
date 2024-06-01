package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "strconv"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

    for {
        fmt.Fprint(os.Stdout, "$ ")

        // Wait for user input
        input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

        fields := strings.Fields(input)
        if fields[0] == "exit" {
            exitCode, _ := strconv.Atoi(fields[1])
            os.Exit(exitCode)
        }
            

        fmt.Printf("%s: command not found\n", strings.Trim(input, "\n"))
    }
}
