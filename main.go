package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	nArgs := len(argsWithoutProg)
	var base_url string
	if nArgs < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if nArgs > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		base_url = argsWithoutProg[0]
		fmt.Println("starting crawl of: %s", base_url)

		resp, _ := http.Get(base_url)
		bytes, _ := io.ReadAll(resp.Body)
		html := string(bytes)
		urls, _ := getURLsFromHTML(html, base_url)
		fmt.Println(urls)
	}

}
