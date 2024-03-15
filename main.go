package main

const DEFAULT_URL = "https://pokeapi.co/api/v2/location-area"

func main() {
	cfg := config{
		nextLocationAreaURL: DEFAULT_URL,
	}
	startRepl(&cfg)
}
