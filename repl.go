package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Kriss-Kolak/pokedexcli/internal/pokeapi"
	"github.com/peterh/liner"
)

var errExit = errors.New("closing")

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
		"catch": {
			name:        "catch",
			description: "Catch pokemon",
			callback:    pokeapi.CatchPokemon,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect owned pokemon",
			callback:    pokeapi.InspectPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Get list of all caught pokemons",
			callback:    pokeapi.Pokedex,
		},
	}
}
func repl() {
	config := pokeapi.GetConfig()
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	for {
		input, err := line.Prompt("Pokedex > ")
		if err == liner.ErrPromptAborted {
			continue
		}
		if err != nil {
			return
		}
		line.AppendHistory(input)
		cleaned := cleanInput(input)
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
			if errors.Is(err, errExit) {
				break
			}
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
