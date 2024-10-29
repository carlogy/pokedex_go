package pokeapi

import (
	"testing"
	"time"
)

func TestCreatePokedex(t *testing.T) {
	pokedex := NewPokedex()

	if pokedex.PokedexEntries == nil {
		t.Error("pokedex is nil")
	}

}

func TestAddPokemon(t *testing.T) {
	c := NewClient(1 * time.Minute)

	pokemon := PokemonEntry{
		name:   "chardmander",
		height: 6,
		weight: 85,
		stats: []statType{
			{
				name:      "hp",
				base_stat: 39,
			}, {
				name:      "attack",
				base_stat: 52,
			},
		},
		types: []string{"fire"},
	}

	c.pokedex.Add(pokemon)
	_, ok := c.pokedex.PokedexEntries[pokemon.name]
	if !ok {
		t.Errorf("%s not found", pokemon.name)
	}

}

func TestConvertPokemonResponseToPokedexEntry(t *testing.T) {
	c := NewClient(1 * time.Minute)

	pikachu := "pikachu"
	pokemonJSONResponse, err := c.PokemonDetails(&pikachu)

	if err != nil {
		t.Error(err)
	}

	pokemon := c.pokedex.Convert(pokemonJSONResponse)

	if pokemon.name != pikachu {
		t.Errorf("%s doesn't match %s", pokemon.name, pikachu)
	}

}
