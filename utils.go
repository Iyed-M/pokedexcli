package main

import (
	"bufio"
	"fmt"
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

func cleanInput() (cmdName string, arg string, err error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	input := scanner.Text()

	if err := scanner.Err(); err != nil {
		return "", "", err
	}
	if input == "" {
		return "", "", nil
	}

	cmd, arg, err := parseInput(input)
	if err != nil {
		return "", "", err
	}

	fmt.Println("cmd : ", cmd)
	fmt.Println("arg : ", arg)
	return strings.ToLower(cmd), arg, nil
}

func parseInput(input string) (cmd string, arg string, err error) {
	cmdAndArgs := strings.Split(input, " ")
	if len(cmdAndArgs) == 2 {
		return cmdAndArgs[0], cmdAndArgs[1], nil
	}
	if len(cmdAndArgs) == 1 {
		return cmdAndArgs[0], "", nil
	}
	return "", "", fmt.Errorf("should have only one arg for got %v", len(cmdAndArgs)-1)
}
