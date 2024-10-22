package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Type the name of pokemon you wish to get more details on. Enter 'quit' to exit the program",
			// callback: commandHelp,
		},
		"quit": {
			name:        "quit",
			description: "Quits the Pokedex",
			// callback: commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		cleanedText := cleanInput(text)

		if len(cleanedText) == 0 {
			continue
		}
		if len(cleanedText) == 1 {
			switch cleanedText[0] {
			case "help":
				fmt.Println(commands["help"].description)
			case "quit":
				fmt.Println("...quiting Pokedex")
				return
			default:
				fmt.Println("Echoing: ", cleanedText)
			}
		} else {
			fmt.Println("Echoing: ", cleanedText)
		}

	}

}

func cleanInput(str string) []string {

	allLower := strings.ToLower(str)
	words := strings.Fields(allLower)

	return words
}
