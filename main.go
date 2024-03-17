package main

import (
	"time"

	"github.com/Iyed-M/pokedexcli/api"
)

const DefaultLocaionsURL = "https://pokeapi.co/api/v2/location-area"

func main() {
	cfg := config{
		nextLocationAreaURL: DefaultLocaionsURL,
		client:              api.NewClient(100*time.Second, 5*time.Second),
	}
	startRepl(&cfg)
}
