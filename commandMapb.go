package main

import (
	"fmt"
)

func commandMapb(cfg *config) error {
	URL := cfg.previousLocationAreaURL
	if URL == "" {
		fmt.Println("noPreviousMaps..\n", "Use map to go forward before going backwards\n", "use help for more info")
		return nil
	}
	locationAreaResp, err := cfg.client.GetLocations(URL)
	if err != nil {
		return err
	}

	for _, location := range locationAreaResp.Results {
		fmt.Printf(" - %s\n", location.Name)
	}
	cfg.nextLocationAreaURL = locationAreaResp.Next
	cfg.previousLocationAreaURL = locationAreaResp.Previous

	return nil
}
