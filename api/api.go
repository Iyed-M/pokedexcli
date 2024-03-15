package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetLocations(URL string) (areasResponse locationAreas, err error) {
	response, err := http.Get(URL)

	if err != nil {
		return locationAreas{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return locationAreas{}, err
	}

	areasData := locationAreas{}
	if err = json.Unmarshal(body, &areasData); err != nil {
		return locationAreas{}, err
	}

	return areasData, nil
}
