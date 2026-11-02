package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Usage: explore <area-name>")
	}
	resp, err := cfg.clientAPI.ExploreLocation(args[0])
	if err != nil {
		return err
	}
	for _, encounter := range resp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
