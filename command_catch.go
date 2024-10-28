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
	fmt.Println("Pokemon in Area: ")
	fmt.Println(res)

	return nil
}
