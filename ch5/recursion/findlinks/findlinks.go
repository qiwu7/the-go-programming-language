package findlinks

import (
	"fmt"

	"golang.org/x/net/html"
)

func Findlinks1(n *html.Node) {
	for _, link := range visit(nil, n) {
		fmt.Println(link)
	}
}

func Findlinks2(n *html.Node) {
	for _, link := range visitRecursively(nil, n) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// ex5.1
// traverse the n.FirstChild linked list using recursive calls to visit
// instead of a loop
func visitRecursively(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visitRecursively(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visitRecursively(links, n.NextSibling)
	}
	return links
}
