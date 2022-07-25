package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"

	"57/the-go-programming-language/ch5/recursion/findlinks"
	"57/the-go-programming-language/ch5/recursion/outline"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	findlinks.Finklinks1(doc)
	outline.Outline(nil, doc)
}
