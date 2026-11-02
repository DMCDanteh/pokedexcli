package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, _ ...string) error {
	resp, err := cfg.clientAPI.ListLocations(cfg.nextPageURL)
	if err != nil {
		return err
	}
	cfg.nextPageURL = resp.Next
	cfg.previousPageURL = resp.Previous
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, _ ...string) error {
	if cfg.previousPageURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.clientAPI.ListLocations(cfg.previousPageURL)
	if err != nil {
		return err
	}
	cfg.nextPageURL = resp.Next
	cfg.previousPageURL = resp.Previous
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
