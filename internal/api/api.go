package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jacksterdealeo/pokedex/internal/pokecache"
)

func GetAPIResponse(url string, cache *pokecache.Cache) ([]byte, error) {
	if data, found := cache.Get(url); found {
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("Response failed with status code: %d\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return []byte{}, err
	}
	cache.Add(url, body)
	return body, nil
}
