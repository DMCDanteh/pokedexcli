package pokeapi

import (
	"net/http"
	pokecache "pokedexcli/internal"
	"time"
)

type Client struct {
	client http.Client
	cache  *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		client: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5 * time.Second),
	}
}
