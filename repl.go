package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	nextURL *string
	prevURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	trimText := strings.ToLower(strings.TrimSpace(text))
	splitText := strings.Split(trimText, " ")
	return splitText
}

func getCliCommands() map[string]cliCommand {
	cliCommands := map[string]cliCommand{
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
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
	}
	return cliCommands
}

func startRepl(cfg *config) {
	cliCommands := getCliCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		} else {
			cmd := cleaned[0]
			cliCmd, ok := cliCommands[cmd]
			if ok {
				err := cliCmd.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
