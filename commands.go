package main

import (
	"fmt"
	"os"

	"github.com/alyssajjc/pokedexcli/internal/pokeapi"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	cliCommands := getCliCommands()
	for _, cmd := range cliCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locationAreaList, err := pokeapi.GetLocationAreas(cfg.nextURL)
	if err != nil {
		return err
	}
	cfg.nextURL = locationAreaList.Next
	cfg.prevURL = locationAreaList.Prev
	locations := locationAreaList.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locationAreaList, err := pokeapi.GetLocationAreas(cfg.prevURL)
	if err != nil {
		return err
	}
	cfg.nextURL = locationAreaList.Next
	cfg.prevURL = locationAreaList.Prev
	locations := locationAreaList.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}
