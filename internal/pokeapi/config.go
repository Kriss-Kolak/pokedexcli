package pokeapi

import (
	"net/http"
	"time"

	"github.com/Kriss-Kolak/pokedexcli/internal/pokecache"
)

const baseURL string = "https://pokeapi.co/api/v2/location-area"

type Config struct {
	Client   *http.Client
	Cache    *pokecache.Cache
	Next     string
	Previous string
}

func GetConfig() *Config {
	return &Config{
		Client:   GetClient(),
		Cache:    pokecache.NewCache(time.Duration(60) * time.Second),
		Next:     "",
		Previous: ""}
}
