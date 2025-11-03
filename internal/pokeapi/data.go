package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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
