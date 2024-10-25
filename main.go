package main

import pokeapi "github.com/carlogy/pokedex_go/internal/pokeAPI"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationListURL *string
	prevLocationListURL *string
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
