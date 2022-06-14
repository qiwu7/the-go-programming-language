package main

import(
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := map[string]int{}
	args := os.Args[1:]
	if len(args) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)	
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
		}
	}
	
	for item, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, item)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
	}
}
