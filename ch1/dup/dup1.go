package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := map[string]int{}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if strings.ToLower(text) == "exit" {
			break
		}
		counts[text]++
	}
	
	for item, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, item)
		}
	}
}

