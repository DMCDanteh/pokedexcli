package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(areaName string) (LocationAreaDetailResponse, error) {
	fullURL := baseURL + "/location-area/" + areaName
	body, ok := c.cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return LocationAreaDetailResponse{}, err
		}
		resp, err := c.client.Do(req)
		if err != nil {
			return LocationAreaDetailResponse{}, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return LocationAreaDetailResponse{}, fmt.Errorf(
				"API request failed with status: %d", resp.StatusCode,
			)
		}
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaDetailResponse{}, err
		}
		c.cache.Add(fullURL, body)
	}
	locationAreaDetailResponse := LocationAreaDetailResponse{}
	err := json.Unmarshal(body, &locationAreaDetailResponse)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}
	return locationAreaDetailResponse, nil
}
