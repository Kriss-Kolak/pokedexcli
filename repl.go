package main

import "strings"

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

var commandMap = map[string]cliCommand{
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
