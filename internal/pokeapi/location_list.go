package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	fullURL := baseURL
	if pageURL == nil {
		fullURL += "/location-area"
	} else {
		fullURL = *pageURL
	}
	body, ok := c.cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		resp, err := c.client.Do(req)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return LocationAreaResponse{}, fmt.Errorf(
				"API request failed with status: %d", resp.StatusCode,
			)
		}
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		c.cache.Add(fullURL, body)
	}

	locationAreaResp := LocationAreaResponse{}
	err := json.Unmarshal(body, &locationAreaResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	return locationAreaResp, nil
}
