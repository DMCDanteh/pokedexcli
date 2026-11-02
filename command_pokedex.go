package main

import "fmt"

func commandPokedex(cfg *config, _ ...string) error {
	if len(cfg.pokedex) == 0 {
		return fmt.Errorf("You have not caught any pokemon yet.")
	}
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
