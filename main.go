package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	nArgs := len(argsWithoutProg)
	var base_url string
	if nArgs < 3 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	} else if nArgs > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		base_url = argsWithoutProg[0]
		maxConcurrency,err := strconv.Atoi(argsWithoutProg[1])
		if err!=nil{
			fmt.Printf("Error convering: %s to int",argsWithoutProg[1])
		}

		maxPages,err := strconv.Atoi(argsWithoutProg[2])
		if err!=nil{
			fmt.Printf("Error convering: %s to int",argsWithoutProg[2])
		}
		// base_url := "https://en.wikipedia.org/wiki/Main_Page"
		// pages := make(map[string]int)
		// crawlPage(base_url, base_url, pages)

		cfg, err := configure(base_url, maxConcurrency, maxPages)
		if err != nil {
			fmt.Printf("Error configuring: %v\n", err)
		}

		cfg.wg.Add(1)
		go cfg.crawlPage(base_url)
		cfg.wg.Wait()
		fmt.Printf("Done crawling.\n")

		pages := cfg.pages
		printReport(pages,base_url)

	}
}
