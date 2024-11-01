# Pokedex

## Cloning Repo

Run `git clone https://github.com/carlogy/pokedex_go.git`

## Make file

* Running application
  - Run `make run` this will execute go build and run the compiled binary
* Tests
  - Verbose test results run `make test-v`
  - Regular test results run `make test`
* Compilation
  - Cross-compile run `make compile`
* All command
  - Running `make all` will test then cross compile

## Pokedex Commands

This is a breakdown of the current commands supported by the pokedex.

* `help` - similar to man, will list all inputs and a brief description
* `map` - makes a paginated call to the PokeAPI and returns 20 locations
  - Make subsequent calls for 20 more locations
* `mapb` - returns the 20 previous locations
* `explore {location-name-here}` - explores that current location and returns Pokemon in that area
* `catch {pokemon name}` - attemps to capture the specified pokemon.
  - if captured the pokemon is stored in your pokedex
* `inspect {pokemon name}` - inspects a given pokemon that you caught and prints specific data on that pokemon.
* `pokedex` - prints out all of your pokedex entries.
* `quit` - quits the program
