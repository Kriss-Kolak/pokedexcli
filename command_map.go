package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Count    int
	Next     string
	Previous string
	Results  []Results
}

type Results struct {
	Name string
	Url  string
}

func commandMap(config *Config) error {
	var url string

	if config.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = config.Next
	}

	res, err := http.Get(url)
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
