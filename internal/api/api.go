package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetAPIResponse(url string) ([]byte, error) {
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
	return body, nil
}
