package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	shell := currentShell()
	switch shell {
	case "bash":
		fmt.Println("bash")
	case "zsh":
		fmt.Println("zsh")
	default:
		fmt.Println("Unknown or unsupported shell. Exiting.")
		os.Exit(1)
	}
}

func currentShell() string {
	shell := os.Getenv("SHELL")
	shellSlice := strings.Split(shell, "/")
	shell = shellSlice[len(shellSlice)-1]
	return shell
}
