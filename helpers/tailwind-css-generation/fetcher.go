package tailwindcssgeneration

import (
	"fmt"
	"io"
	"net/http"
)

// GETFetcher performs a simple HTTP GET request and returns the response body
// as a string. The function is agnostic to the content type, usage context,
// or source of the data. It performs no caching and no retries.
func GETFetcher(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("http get failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("http get returned status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("http get read failed: %w", err)
	}

	return string(data), nil
}
