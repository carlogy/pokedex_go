package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if args == nil || args[0] == "" {
		return errors.New("Inspect command without pokemon name entered")
	}

	name := args[0]

	pokemon, err := cfg.pokeapiClient.Inspect(name)

	if err != nil {
		return err
	}

	fmt.Println(pokemon)

	return nil
}
