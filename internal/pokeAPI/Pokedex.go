package pokeapi

type Pokedex struct {
	pokedexEntries map[string]PokemonEntry
}

type PokemonEntry struct {
	name   string
	height int
	weight int
	stats  stat
	types  []string
}

type stat struct {
	name      string
	base_stat int
}

// to do create Pokedex entry and add to map
func (pokedex *Pokedex) Add(name string, pokemon Pokemon) {

}
