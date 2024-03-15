package main

import (
	"bufio"
	"os"
	"strings"
)

func getCmdFromStdin(commands map[string]clicommand) (clicommand, error) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input := scanner.Text()
	if err := scanner.Err(); err != nil {
		return clicommand{}, err
	}
	if input == "" {
		return clicommand{}, nil
	}
	command, ok := commands[input]
	if !ok {
		return commands["help"], nil
	}
	return command, nil

}

func cleanInput() (cmdName string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if err := scanner.Err(); err != nil {
		return "", err
	}
	if input == "" {
		return "", nil
	}
	return strings.ToLower(input), nil
}
