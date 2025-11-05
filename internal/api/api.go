package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokeAPIResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetPokeAPIResponse(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

// Do not use this function if you plan on caching bytes.
func GetPokeAPIResponseMarshaled(url string) (*PokeAPIResponse, error) {
	body, err := GetPokeAPIResponse(url)
	if err != nil {
		return nil, err
	}
	var res *PokeAPIResponse
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, fmt.Errorf("Couldn't Unmarshal json body\nerr: %v\njson: %v", err, body)
	}
	return res, nil
}
