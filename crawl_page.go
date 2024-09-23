package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// SKIP IF RUNNING OTHER WEBSITE
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}
	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}
	// normalize & check current url
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error normalizing url %s: %v", normalizedURL, err)
	}

	// update pages
	value, found := pages[normalizedURL]
	if found {
		pages[normalizedURL] = value + 1
		return
	} else {
		pages[normalizedURL] = 1
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)
	// get html, find urls
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing url %s: %v\n", normalizedURL, err)
	}
	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("Error getting urls from url %s: %v\n", normalizedURL, err)
	}

	// crawl for all found
	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
