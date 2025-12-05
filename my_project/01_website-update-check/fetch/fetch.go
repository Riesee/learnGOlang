package fetch

import (
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := io.ReadAll(resp.Body)
	if err != nil {
		return doc, err
	}
	return doc, nil
}