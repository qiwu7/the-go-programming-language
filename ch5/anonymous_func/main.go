package main

import (
	"57/the-go-programming-language/ch5/anonymous_func/crawl"
	"os"
)

func main() {
	crawl.BFS(crawl.Crawl, os.Args[1:])
}
