package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
)

// get data from location-area endpoint (used in forward and backward map)
func GetData(config *Config, url string) error {

	body, ok := config.Cache.Get(url)
	if !ok {
		res, err := config.Client.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		config.Cache.Add(url, body)
	}

	var location Location
	if err := json.Unmarshal(body, &location); err != nil {
		return err
	}

	config.Next = location.Next
	config.Previous = location.Previous

	for _, result := range location.Results {
		fmt.Println(result.Name)
	}

	return nil
}

// Search forward with config.Next url
func CommandMapF(config *Config, argument string) error {
	var err error
	if config.Next == "" {
		err = GetData(config, baseURL)
	} else {
		err = GetData(config, config.Next)
	}
	return err
}

// Search backward with config.Previous url
func CommandMapB(config *Config, argument string) error {
	var err error
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		err = GetData(config, config.Previous)
	}
	return err
}

// Get available pokemon list from given location
func GetPokemonList(config *Config, locationName string) error {
	url := baseURL + "/" + locationName
	body, ok := config.Cache.Get(url)
	if !ok {
		res, err := config.Client.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		config.Cache.Add(url, body)
	}
	var locationID LocationAreaID
	if err := json.Unmarshal(body, &locationID); err != nil {
		return err
	}
	fmt.Printf("Exploring %s\n", locationID.Name)
	fmt.Println("Found Pokemon:")
	for _, element := range locationID.PokemonEncounters {
		fmt.Printf("- %s \n", element.Pokemon.Name)
	}

	return nil
}

// Approach to catch a pokemon
func CatchPokemon(config *Config, pokemonName string) error {
	url := pokemonURL + "/" + pokemonName

	res, err := config.Client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	throw := rand.Float32()
	if throw >= 0.5 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.Pokemons[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

// Inspect basic stats about pokemon
func InspectPokemon(config *Config, pokemonName string) error {
	pokemon, ok := config.Pokemons[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemon.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}

	return nil
}

// Get list of all caught pokemons
func Pokedex(config *Config, argument string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range config.Pokemons {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
