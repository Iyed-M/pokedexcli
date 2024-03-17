package main

import "github.com/Iyed-M/pokedexcli/api"

type clicommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	nextLocationAreaURL     string
	previousLocationAreaURL string
	client                  *api.Client
}

func getAvailableCommands() map[string]clicommand {
	return map[string]clicommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},

		"map": {
			name:        "map",
			description: "Get 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "kmap",
			description: "Get previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback:    commandCatch,
		},
	}
}
