package main

import (
	"os"
	"path/filepath"
	"strings"
)

var pathDirectories []string

func init() {
	pathDirectories = filepath.SplitList(os.Getenv("PATH"))
}

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

func resolveFromPathVariable(command string) (path string, found bool) {
	for _, path := range pathDirectories {
		if fileInPath(path, command) {
			return path, true
		}
	}

	return "", false
}
