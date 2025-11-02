package pokeapi

import (
	"net/http"
	"time"
)

func get_client() *http.Client {
	c := http.Client{Timeout: time.Duration(5) * time.Second}
	return &c
}
