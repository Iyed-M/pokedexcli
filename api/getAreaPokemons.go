package api

import (
	"encoding/json"
	"io"
)

func (c *Client) GetAreaPokemons(areaName string) (AreaPokemons, error) {
	var unparsedData []byte
	pokemons := AreaPokemons{}
	URL := "https://pokeapi.co/api/v2/location-area/" + areaName

	cachedData, ok := c.pokeCache.Get(URL)
	// cache hit
	if ok {
		unparsedData = cachedData
	} else {
		// chache miss
		resp, err := c.httpClient.Get(URL)
		if err != nil {
			return AreaPokemons{}, err
		}

		defer resp.Body.Close()
		unparsedData, err = io.ReadAll(resp.Body)
		if err != nil {
			return AreaPokemons{}, err
		}
		// add to cache
		c.pokeCache.Add(URL, unparsedData)
	}

	if err := json.Unmarshal(unparsedData, &pokemons); err != nil {
		return AreaPokemons{}, err
	}
	return pokemons, nil
}
