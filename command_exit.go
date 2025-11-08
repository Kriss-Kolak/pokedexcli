package main

import (
	"fmt"

	"github.com/Kriss-Kolak/pokedexcli/internal/pokeapi"
)

func commandExit(config *pokeapi.Config, argument string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return errExit
}
