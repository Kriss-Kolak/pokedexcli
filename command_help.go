package main

import (
	"fmt"

	"github.com/Kriss-Kolak/pokedexcli/internal/pokeapi"
)

func commandHelp(config *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, value := range getCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
