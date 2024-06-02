package main

import (
	"os"
	"path/filepath"
	"strings"
    "errors"
)

func resolveHomePath(targetPath string) (newPath string, mapped bool) {
	topComponent, rest, found := strings.Cut(targetPath, "/")

	if topComponent != "~" {
		return targetPath, false
	}

	homePath := os.Getenv("HOME")
	if !found {
		return homePath, true
	}

	return filepath.Join(homePath, rest), true
}

func resolvePath(args string) string {
	if filepath.IsAbs(args) {
		return args
	}

	mappedHomePath, isHome := resolveHomePath(args)
	if isHome {
		return mappedHomePath
	}

	return filepath.Join(os.Getenv("PWD"), args)
}

func findExecutableFile(command string) (path string, found bool) {
	commandPath := resolvePath(command)
	if _, err := os.Stat(commandPath); errors.Is(err, os.ErrNotExist) {
		if path, found := findInPathVariable(command); found {
			commandPath = filepath.Join(path, command)
			return commandPath, true
		}

        return "", false
	}

    return commandPath, true
}

func findInDirectory(path string, command string) (found bool) {
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

func findInPathVariable(command string) (path string, found bool) {
	pathDirectories := filepath.SplitList(os.Getenv("PATH"))
	for _, path := range pathDirectories {
		if findInDirectory(path, command) {
			return path, true
		}
	}

	return "", false
}
