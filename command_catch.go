package main

import (
	"errors"
	"fmt"
)

func callbackCatch(cfg *config, args ...string) error {
	if args == nil || args[0] == "" {
		return errors.New("Explore command without location entered, please include location to explore")
	}
	location := args[0]

	res, err := cfg.pokeapiClient.PokemonDetails(&location)
	if err != nil {
		return err
	}

	attemptCatch, err := cfg.pokeapiClient.CatchPokemon(res)
	if err != nil {
		return err
	}

	if !attemptCatch {
		fmt.Printf("%s escaped!\n", res.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", res.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}
