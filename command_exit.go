package main

import (
	"fmt"
	"os"

	"github.com/Kriss-Kolak/pokedexcli/internal/pokeapi"
)

func commandExit(config *pokeapi.Config, argument string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
