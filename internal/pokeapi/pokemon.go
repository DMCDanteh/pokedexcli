package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Pokemon(pokemon string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + pokemon
	body, ok := c.cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return Pokemon{}, err
		}
		resp, err := c.client.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return Pokemon{}, fmt.Errorf(
				"API request failed with status: %d", resp.StatusCode,
			)
		}
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(fullURL, body)
	}
	pokemonResponse := Pokemon{}
	err := json.Unmarshal(body, &pokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemonResponse, nil
}
