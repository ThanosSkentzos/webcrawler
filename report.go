package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("==============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("==============================")

	keys := make([]string, len(pages))

	i := 0
	for k := range pages {
		keys[i] = k
		i++
	}
	sort.SliceStable(keys,
		func(i, j int) bool {
		return pages[keys[i]] > pages[keys[j]]
		})

	for _, url := range keys {
		count := pages[url]
		fmt.Printf("Found %d internal links to %s\n", count, url)
	}

}
