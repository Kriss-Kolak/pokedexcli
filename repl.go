package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {

	//Lower the input
	lowered := strings.ToLower(text)
	//Split the input by white spaces
	splited := strings.Fields(lowered)

	return splited
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
func repl() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		userInput := s.Text()
		cleaned := cleanInput(userInput)
		var command string = ""
		if len(cleaned) == 0 {
			continue
		}
		command = cleaned[0]
		fc, ok := getCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := fc.callback()
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
