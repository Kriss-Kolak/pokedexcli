package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kriss-Kolak/pokedexcli/internal/pokeapi"
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
	callback    func(*pokeapi.Config, string) error
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
		"map": {
			name:        "map",
			description: "Display next 20 locations",
			callback:    pokeapi.CommandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations",
			callback:    pokeapi.CommandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore location for pokemons",
			callback:    pokeapi.GetPokemonList,
		},
	}
}
func repl() {
	config := pokeapi.GetConfig()
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
		var argument string = ""
		if len(cleaned) > 1 {
			argument = cleaned[1]
		}

		fc, ok := getCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := fc.callback(config, argument)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
