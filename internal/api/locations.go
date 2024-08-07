package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (PokemonLocations, error) {
	url := "https://pokeapi.co/api/v2/" + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonLocations{}, err
	}

	locationsResp := PokemonLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return PokemonLocations{}, err
	}

	return locationsResp, nil
}
