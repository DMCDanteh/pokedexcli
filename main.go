package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	cfg := config{
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Displays all locations on the map",
				callback:    commandMapf,
			},
			"mapb": {
				name:        "mapb",
				description: "Diplays all locations on the previous map",
				callback:    commandMapb,
			},
			"explore": {
				name:        "explore",
				description: "Explore a location area and list its Pokemon",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Attempt to catch a Pokemon",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect a caught Pokemon",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "Displays all your caught Pokemon",
				callback:    commandPokedex,
			},
		},
		clientAPI: pokeapi.NewClient(5 * time.Second),
		pokedex:   map[string]pokeapi.Pokemon{},
	}
	startRepl(&cfg)
}
