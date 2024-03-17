package main

import (
	"fmt"
)

func commandMap(cfg *config, _ string) error {
	URL := cfg.nextLocationAreaURL
	if URL == "" {
		fmt.Println("== No more locations", "== Resetting...")
		URL = DefaultLocaionsURL
	}
	// get location from api
	areasResp, err := cfg.client.GetLocations(URL)
	if err != nil {
		return err
	}

	// list locations from response
	for _, location := range areasResp.Results {
		fmt.Printf(" - %s\n", location.Name)
	}

	cfg.nextLocationAreaURL = areasResp.Next
	cfg.previousLocationAreaURL = areasResp.Previous
	return nil
}
