package main

import (
	"fmt"
	"sort"
)

func main() {
	// argsWithoutProg := os.Args[1:]
	// nArgs := len(argsWithoutProg)
	// var base_url string
	// if nArgs < 1 {
	// 	fmt.Println("no website provided")
	// 	os.Exit(1)
	// } else if nArgs > 1 {
	// 	fmt.Println("too many arguments provided")
	// 	os.Exit(1)
	// } else {
	// 	base_url = argsWithoutProg[0]
	base_url := "https://google.com"
	pages := make(map[string]int)

	crawlPage(base_url, base_url, pages)

	keys := make([]string, len(pages))

	i := 0
	for k := range pages {
		keys[i] = k
		i++
	}
	sort.SliceStable(keys,
		func(i, j int) bool {
			return pages[keys[i]] < pages[keys[j]]
		})

	for _, url := range keys {
		count := pages[url]
		fmt.Printf("%d - %s\n", count, url)
	}

	// }
}
