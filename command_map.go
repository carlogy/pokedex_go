package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationListURL)
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
