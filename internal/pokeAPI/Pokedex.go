package pokeapi

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
