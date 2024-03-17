package main

import "fmt"

func commandHelp(_ *config, _ string) error {
	fmt.Println("Usage :")
	fmt.Println()
	m := getAvailableCommands()

	for cmd := range m {
		fmt.Printf("\t%s : ", cmd)
		fmt.Printf("%s\n", m[cmd].description)
	}

	return nil
}
