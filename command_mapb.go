package main

import (
	"errors"
	"fmt"
)

func callbackMapb(cfg *config) error {

	if cfg.prevLocationListURL == nil {

		return errors.New("Already provided the first page of locations list")

	}
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationListURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")

	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationListURL = res.Next
	cfg.prevLocationListURL = res.Previous

	return nil
}
