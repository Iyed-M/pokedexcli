package main

import (
	"time"

	"github.com/Iyed-M/pokedexcli/api"
)

<<<<<<< Updated upstream
const DefaultLocaionsURL = "https://pokeapi.co/api/v2/location-area"
=======
const (
	DefaultLocaionsURL   = "https://pokeapi.co/api/v2/location-area"
	CapturedPokemonsFile = "captured_pokemons.json"
)
>>>>>>> Stashed changes

func main() {
	cfg := config{
		nextLocationAreaURL: DefaultLocaionsURL,
		client:              api.NewClient(100*time.Second, 5*time.Second),
<<<<<<< Updated upstream
=======
		pokemons:            loadPokemons(),
>>>>>>> Stashed changes
	}
	startRepl(&cfg)
}
