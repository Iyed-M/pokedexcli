package main

import (
	"time"

	"github.com/Iyed-M/pokedexcli/api"
)

const DEFAULT_URL = "https://pokeapi.co/api/v2/location-area"

func main() {
	cfg := config{
		nextLocationAreaURL: DEFAULT_URL,
		client:              api.NewClient(100*time.Second, 5*time.Second),
	}
	startRepl(&cfg)
}
