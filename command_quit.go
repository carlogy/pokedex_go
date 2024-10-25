package main

import (
	"fmt"
	"os"
)

func callbackQuit(cfg *config) error {
	fmt.Println("Quiting Pokedex...")
	os.Exit(0)

	return nil
}
