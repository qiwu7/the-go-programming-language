package main

import (
	"57/the-go-programming-language/ch5/function_values/outline"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("url %s: %v\n", url, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("getting %s: %s\n", url, resp.Status)
			continue
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Printf("parsing %s as HTML: %v\n", url, err)
		}

		outline.Outline(doc)
	}
}
