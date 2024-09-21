package main

import (
	"fmt"
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
		// fmt.Printf("starting crawl of: %s\n", base_url)

		// resp, err := http.Get(base_url)
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		// defer resp.Body.Close()
		// bytes, _ := io.ReadAll(resp.Body)
		// html := string(bytes)
		// urls, _ := getURLsFromHTML(html, base_url)
		// fmt.Println(urls)

		// getHTML test
		html, err := getHTML(base_url)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Print("Got:\n", html)
		}
	}

}
