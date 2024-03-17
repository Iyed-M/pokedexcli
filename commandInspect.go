package main

import "fmt"

func commandInspect(cfg *config, pokemonName string) error {
	pokemon, ok := cfg.pokemons[pokemonName]
	if !ok {
		fmt.Printf("You have'nt caught %s yet", pokemonName)
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	return nil
}
