package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Iyed-M/pokedexcli/pokecache"
)

type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
}

func NewClient(interval time.Duration, timeout time.Duration) *Client {
	return &Client{
		httpClient: http.Client{Timeout: timeout},
		pokeCache:  pokecache.NewCache(interval),
	}
}

func (c *Client) GetLocations(URL string) (locationAreas, error) {
	cachedData, ok := c.pokeCache.Get(URL)
	var unparesedLocations []byte

	// cach hit
	if ok {
		fmt.Println("== Cache hit")
		unparesedLocations = cachedData
	} else {
		// cache miss
		resp, err := c.httpClient.Get(URL)
		if err != nil {
			return locationAreas{}, err
		}
		defer resp.Body.Close()
		unparesedLocations, err = io.ReadAll(resp.Body)
		if err != nil {
			return locationAreas{}, err
		}
		// store in cache
		c.pokeCache.Add(URL, unparesedLocations)
	}

	// parse locations
	areas := locationAreas{}
	err := json.Unmarshal(unparesedLocations, &areas)
	if err != nil {
		return locationAreas{}, err
	}
	return areas, err
}
