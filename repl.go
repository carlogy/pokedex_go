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
	callback    func() error
}

func startRepl() {

	// commands := map[string]cliCommand{
	// 	"help": {
	// 		name:        "help",
	// 		description: "1. Type the name of pokemon you wish to get more details on\n2. Type 'map' to get 20 locations, for more type 'map' again\n3. Type 'mapb' for the previous 20 locations \n4. Enter 'quit' or 'exit' to exit the program",
	// 		// callback: commandHelp,
	// 	},
	// 	"quit": {
	// 		name:        "quit",
	// 		description: "Quits the Pokedex",
	// 		// callback: commandExit,
	// 	},
	// 	"map": {
	// 		name:        "map",
	// 		description: "displays 20 location areas in the Pokemon world, each subsequent call display 20 more locations.",
	// 	},
	// }

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

		command.callback()

		// switch command {
		// case "help":
		// 	commands["help"].callback()
		// case "quit":
		// 	commands["quit"].callback()
		// case "exit":
		// 	commands["quit"].callback()
		// case "map":
		// 	var locations pokeAPI.Locations

		// 	// url, err := pokeAPI.NewParsedURL("https://pokeapi.co/api/v2/location/")
		// 	// if err != nil {
		// 	// 	fmt.Println("map error: %w", err)
		// 	// }
		// 	// fmt.Println(url)

		// 	locations, err := pokeAPI.GetLocations("https://pokeapi.co/api/v2/location/")
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}

		// 	fmt.Println(locations)

		// default:
		// 	fmt.Println("Echoing: ", cleanedText)
		// }

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
	}
}
