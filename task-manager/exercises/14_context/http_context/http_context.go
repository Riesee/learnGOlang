package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchURL(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	url := "http://localhost:3004/slow"

	fmt.Println("Fetching URL:", url)
	start := time.Now()
	body, err := fetchURL(ctx, url)

	fmt.Println("Took:", time.Since(start))
	if err != nil {
        fmt.Println("❌ Timeout/Hata:", err)
        return
    }
    fmt.Println("✅ Cevap:", body)
}