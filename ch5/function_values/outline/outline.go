package outline

import (
	"fmt"

	"golang.org/x/net/html"
)

func Outline(n *html.Node) {
	forEachNode(n, -1, startElement, endElement)

}

func forEachNode(n *html.Node, depth int, pre, post func(n *html.Node, depth int)) {
	if pre != nil {
		pre(n, depth)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, depth+1, pre, post)
	}

	if post != nil {
		post(n, depth)
	}
}

func startElement(n *html.Node, depth int) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node, depth int) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
