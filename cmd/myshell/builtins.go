package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var builtInHandlers map[string]func(string)

func init() {
	builtInHandlers = map[string]func(string){
		"exit": exitCommand,
		"echo": echoCommand,
		"type": typeCommand,
		"pwd":  pwdCommand,
		"cd":   cdCommand,
	}
}

func handleBuiltIn(command string, args string) (ok bool) {
	handler, ok := builtInHandlers[command]
	if ok {
		handler(args)
	}

	return
}

func exitCommand(args string) {
	exitCode, _ := strconv.Atoi(args)
	os.Exit(exitCode)
}

func echoCommand(args string) {
	fmt.Println(args)
}

func pwdCommand(_ string) {
	fmt.Println(os.Getenv("PWD"))
}

func cdCommand(args string) {
	targetPath := resolvePath(args)

	if _, err := os.Stat(targetPath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s: No such file or directory\n", targetPath)
		return
	}

	os.Chdir(targetPath)
	os.Setenv("PWD", targetPath)
}

func typeCommand(args string) {
	if _, ok := builtInHandlers[args]; ok {
		fmt.Printf("%s is a shell builtin\n", args)
	} else if path, found := findExecutableFile(args); found {
		fmt.Printf("%s is %s\n", args, path)
	} else {
		fmt.Printf("%s not found\n", args)
	}
}
