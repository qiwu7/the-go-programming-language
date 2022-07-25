package outline

import (
	"fmt"

	"golang.org/x/net/html"
)

func Outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Outline(stack, c)
	}
}
