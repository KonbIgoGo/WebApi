package handlers

import (
	"io"
	"net/http"
	"time"
)

func readRaw(url string) (io.ReadCloser, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Referer", "https://example.com/")
	req.Header.Set("Cookie", "sessionid=abcd1234; other=xyz")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
