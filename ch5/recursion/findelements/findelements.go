package findelements

import (
	"fmt"

	"golang.org/x/net/html"
)

func FindElements(n *html.Node) {
	counts := map[string]int{}
	visit(counts, n)
	for k, v := range counts {
		fmt.Printf("type: %10s, count: %5d\n", k, v)
	}
}

func visit(m map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(m, c)
	}
}
