package pokeapi

import (
	"errors"
	"fmt"
)

type Pokedex struct {
	PokedexEntries map[string]PokemonEntry
}

type PokemonEntry struct {
	name   string
	height int
	weight int
	stats  []statType
	types  []string
}

type statType struct {
	name      string
	base_stat int
}

func NewPokedex() Pokedex {
	return Pokedex{
		PokedexEntries: make(map[string]PokemonEntry),
	}
}

func (pokedex *Pokedex) Convert(pokemon Pokemon) PokemonEntry {
	Pokemon := PokemonEntry{
		stats:  make([]statType, len(pokemon.Stats)),
		types:  make([]string, len(pokemon.Types)),
		name:   pokemon.Name,
		height: pokemon.Height,
		weight: pokemon.Weight,
	}

	for i, stat := range pokemon.Stats {
		s := statType{name: stat.Stat.Name, base_stat: stat.BaseStat}
		Pokemon.stats[i] = s
	}

	for i, t := range pokemon.Types {
		Pokemon.types[i] = t.Type.Name
	}

	return Pokemon
}

func (pokedex *Pokedex) Add(pokemon PokemonEntry) {

	pokedex.PokedexEntries[pokemon.name] = pokemon
}

func (pokedex *Pokedex) Get(name string) (PokemonEntry, bool) {

	pokemon, ok := pokedex.PokedexEntries[name]
	if !ok {
		return PokemonEntry{}, ok
	}

	return pokemon, ok

}

func (pokedex *Pokedex) GetEntries(p Pokedex) ([]string, error){

	lengthPokedex := len(p.PokedexEntries)

	if lengthPokedex == 0 {
		return []string{}, errors.New("pokedex is empty")
	}

	entries := make([]string, 0, lengthPokedex)

	for i, _ := range p.PokedexEntries{
		entries = append(entries, i)
	}

	return entries, nil
}

func (p PokemonEntry) String() string {

	pokeString := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", p.name, p.height, p.weight)

	for _, stat := range p.stats {
		pokeString += fmt.Sprintf("-%s: %d\n", stat.name, stat.base_stat)
	}

	pokeString += "Types:\n"

	for _, stat := range p.types {
		pokeString += fmt.Sprintf("-%s\n", stat)
	}

	return pokeString
}
