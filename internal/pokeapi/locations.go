package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaList struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

func GetLocationAreas(url *string) (LocationAreaList, error) {
	fullURL := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	if url != nil {
		fullURL = *url
	}
	res, err := http.Get(fullURL)
	if err != nil {
		return LocationAreaList{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return LocationAreaList{}, fmt.Errorf("Failed with status code %d", res.StatusCode)
	}
	if err != nil {
		return LocationAreaList{}, err
	}
	var locationAreaList LocationAreaList
	err = json.Unmarshal(body, &locationAreaList)
	return locationAreaList, err
}
