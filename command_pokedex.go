package main

import "fmt"

func callbackPokedex(cfg *config, args ...string) error {

	fmt.Println("Your Pokedex:")

	entries, err := cfg.pokeapiClient.Pokedex()
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Printf("- %s\n", entry)
	}

	return nil

}
