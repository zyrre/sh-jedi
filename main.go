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
	topTenCommands := make(map[string]int)

	for i := 0; i < 10; i++ {
		var maxKey string
		var maxVal int
		for k, v := range uniqueCommands {
			if v > maxVal {
				maxKey = k
				maxVal = v
			}
		}
		topTenCommands[maxKey] = maxVal
		delete(uniqueCommands, maxKey)
	}
	fmt.Println(topTenCommands)
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
	for i, v := range commands {
		if i < len(commands)-1 {
			uniqueCommands[v+" :: "+commands[i+1]]++
		}
	}
	return uniqueCommands
}
