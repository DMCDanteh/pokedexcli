package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	commands        map[string]cliCommand
	clientAPI       pokeapi.Client
	nextPageURL     *string
	previousPageURL *string
	pokedex         map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		if !scanner.Scan() {
			return
		}
		input := scanner.Text()
		word := cleanInput(input)
		if len(word) == 0 {
			continue
		}
		command, exists := cfg.commands[word[0]]
		if !exists {
			fmt.Print("Unknown command")
			continue
		}
		err := command.callback(cfg, word[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func commandExit(*config, ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, _ ...string) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
