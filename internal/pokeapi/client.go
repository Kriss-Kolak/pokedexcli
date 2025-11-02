package pokeapi

import (
	"net/http"
	"time"
)

func GetClient() *http.Client {
	c := http.Client{Timeout: time.Duration(5) * time.Second}
	return &c
}
