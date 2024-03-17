package main

import "fmt"

func commandExplore(cfg *config, areaName string) error {
	if areaName == "" {
		fmt.Println("Usage : explore <location-area>")
	}
	pokemons, err := cfg.client.GetAreaPokemons(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, pok := range pokemons.PokemonEncounters {
		fmt.Println(" - ", pok.Pokemon.Name)
	}

	return nil
}
