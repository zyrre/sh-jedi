package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	shell := currentShell()
	commands := readHistFile(shell)
	uniqueCommands := countUniqueCommands(commands)
	fmt.Println(uniqueCommands)
}

func currentShell() string {
	shell := os.Getenv("SHELL")
	shellSlice := strings.Split(shell, "/")
	shell = shellSlice[len(shellSlice)-1]
	return shell
}

func readHistFile(shell string) []string {
	var histFile string

	switch shell {
	case "bash":
		histFile = os.Getenv("HOME") + "/.bash_history"
	case "zsh":
		histFile = os.Getenv("HOME") + "/.zsh_history"
	default:
		fmt.Println("Unknown or unsupported shell. Exiting.")
		os.Exit(1)
	}

	file, err := os.Open(histFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	fileString := strings.Split(string(fileBytes), "\n: ")
	for i, v := range fileString {
		fileString[i] = strings.Split(strings.TrimSpace(v), ";")[1]
	}
	return fileString
}

func countUniqueCommands(commands []string) map[string]int {
	uniqueCommands := make(map[string]int)
	for _, v := range commands {
		uniqueCommands[v]++
	}
	return uniqueCommands
}
