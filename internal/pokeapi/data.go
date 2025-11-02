package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func GetData(config *Config, url string) error {
	res, err := config.Client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
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

func CommandMapF(config *Config) error {
	var err error
	if config.Next == "" {
		err = GetData(config, baseURL)
	} else {
		err = GetData(config, config.Next)
	}
	return err
}

func CommandMapB(config *Config) error {
	var err error
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		err = GetData(config, config.Previous)
	}
	return err
}
