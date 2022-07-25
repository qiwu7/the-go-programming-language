package outline

import (
	"fmt"

	"golang.org/x/net/html"
)

var depth int

func ElementByID(n *html.Node, id string) *html.Node {
	return forEachNode2(n, id, startElement2, endElement2)
}

func forEachNode2(
	n *html.Node,
	id string,
	pre, post func(n *html.Node, id string) bool,
) *html.Node {
	if pre != nil && pre(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode2(c, id, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil && post(n, id) {
		return n
	}

	return nil
}

// startElement performs pre actions and returns whether id matches
func startElement2(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
	return matches(n, id)
}

// endElement performs post actions and returns whether id matches
func endElement2(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
	return matches(n, id)
}

func matches(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" {
				if a.Val == id {
					return true
				}
			}
		}
	}
	return false
}
