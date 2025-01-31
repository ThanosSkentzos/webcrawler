package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't get http: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("status over 400 for url %s", rawURL)
	}
	content_type := resp.Header.Get("Content-Type")
	if !strings.Contains(content_type, "text/html") {
		return "", fmt.Errorf("content-type %s is not text/html", resp.Header.Get("content-type"))
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response: %v", err)
	}
	html := string(bytes)

	return html, nil
}
