package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"

	"57/the-go-programming-language/ch5/recursion/findelements"
	"57/the-go-programming-language/ch5/recursion/findlinks"
	"57/the-go-programming-language/ch5/recursion/outline"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "1":
			fmt.Println("Findlinks")
			findlinks.Findlinks1(doc)
		case "2":
			fmt.Println("Outline")
			outline.Outline(nil, doc)
		case "3":
			fmt.Println("Find Elements")
			findelements.FindElements(doc)
		case "4":
			fmt.Println("Find All Links (a, img, script)")
			findlinks.Findlinks3(doc)
		default:
			// do nothing
		}
	}
}
