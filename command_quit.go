package main

import (
	"fmt"
	"os"
)

func callbackQuit() error {
	fmt.Println("Quiting Pokedex...")
	os.Exit(0)

	return nil
}
