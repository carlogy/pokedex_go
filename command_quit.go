package main

import (
	"fmt"
	"os"
)

func callbackQuit(cfg *config, args ...string) error {
	fmt.Println("Quiting Pokedex...")
	os.Exit(0)

	return nil
}
