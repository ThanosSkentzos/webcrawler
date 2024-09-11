package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(badUrl string) (string, error) {
	parsedURL, err := url.Parse(badUrl)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	fullPath := parsedURL.Host + parsedURL.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
