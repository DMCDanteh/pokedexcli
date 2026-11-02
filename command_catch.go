package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Usage: catch <pokemon-name>")
	}
	resp, err := cfg.clientAPI.Pokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)
	catchChance := (1000 - resp.BaseExperience) / 10
	roll := rand.Intn(100)
	if roll < catchChance {
		fmt.Printf("%s was caught!\n", resp.Name)
		fmt.Printf("You may now inspect it with the inspect command.\n")
		cfg.pokedex[resp.Name] = resp
	} else {
		fmt.Printf("%s escaped!\n", resp.Name)
	}
	return nil
}
