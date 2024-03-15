package main

import (
	"fmt"
	"os"
)

func startRepl(cfg *config) {
	commands := getAvailableCommands()

	for {
		fmt.Print("Pokedex > ")
		inputCmd, err := cleanInput()

		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading command : %s\n", err)
		}

		if inputCmd == "" {
			continue
		}
		cmd, ok := commands[inputCmd]
		if !ok {
			fmt.Fprintf(os.Stderr, "command %s not found\n", inputCmd)
			continue
		}
		if err := cmd.callback(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "error executing command %s : %s\n", inputCmd, err)
		}
	}
}
