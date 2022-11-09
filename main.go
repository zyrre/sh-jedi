package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	shell := currentShell()
	fmt.Println(shell)
}

func currentShell() string {
	shell := os.Getenv("SHELL")
	shellSlice := strings.Split(shell, "/")
	shell = shellSlice[len(shellSlice)-1]
	return shell
}
