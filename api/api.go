package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Iyed-M/pokedexcli/pokecache"
)

func GetLocations(URL string) (areasResponse locationAreas, err error) {
	areasJson, err := fetchLocations(URL) // TODO: Add caching
	areasData := locationAreas{}
	if err = json.Unmarshal(areasJson, &areasData); err != nil {
		return locationAreas{}, err
	}
	return areasData, nil
}

func fetchLocations(c *pokecache.Cache, URL string) ([]byte, error) {
	cachedData, ok := c.Get(URL)

	if !ok {
		resp, err := http.Get(URL)
		defer resp.Body.Close()

		if err != nil {
			return nil, err
		}

		locations, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return locations, nil
	}
	return cachedData, nil
}
