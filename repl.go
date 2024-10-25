package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)

		if len(cleanedText) == 0 {
			continue
		}

		commandInput := cleanedText[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandInput]

		if !ok {
			fmt.Println("Invalid command entered")
			continue
		}
		err := command.callback(cfg)

		if err != nil {
			fmt.Println(err)
		}

	}

}

func cleanInput(str string) []string {

	allLower := strings.ToLower(str)
	words := strings.Fields(allLower)

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Provides helpful instructions on how to use the Pokedex",
			callback:    callbackHelp,
		},
		"quit": {
			name:        "quit",
			description: "Quits the Pokedex",
			callback:    callbackQuit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in the Pokemon world, each subsequent call display 20 more locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world, each subsequent call displays the previous 20 locations",
			callback:    callbackMapb,
		},
	}
}
