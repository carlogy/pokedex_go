package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {

	if args == nil || args[0] == "" {
		return errors.New("Explore command without location entered, please include location to explore")
	}
	location := args[0]

	res, err := cfg.pokeapiClient.ListPokemon(&location)
	if err != nil {
		return err
	}
	fmt.Println("Pokemon in Area: ")

	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("- %s\n", *pokemon.Pokemon.Name)
	}
	return nil
}
