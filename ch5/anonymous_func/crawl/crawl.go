package crawl

import (
	"57/the-go-programming-language/ch5/anonymous_func/links"
	"fmt"
	"log"
)

func BFS(f func(item string) []string, worklist []string) {
	visited := map[string]bool{}
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
