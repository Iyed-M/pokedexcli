package main

type clicommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextLocationAreaURL     string
	previousLocationAreaURL string
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
	}
}
