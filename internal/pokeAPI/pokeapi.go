package pokeapi

import (
	"net/http"
	"time"

	pokecache "github.com/carlogy/pokedex_go/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	pokedex    Pokedex
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache:   pokecache.NewCache(cacheInterval),
		pokedex: NewPokedex(),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
