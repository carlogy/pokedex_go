package main

import (
	"fmt"
	"log"

	pokeapi "github.com/carlogy/pokedex_go/internal/pokeAPI"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	res, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas:")

	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	return nil
}
