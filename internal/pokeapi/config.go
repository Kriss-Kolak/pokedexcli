package pokeapi

import "net/http"

const baseURL string = "https://pokeapi.co/api/v2/location-area"

type Config struct {
	Client   *http.Client
	Next     string
	Previous string
}

func GetConfig() *Config {
	return &Config{
		Client:   GetClient(),
		Next:     "",
		Previous: ""}
}
