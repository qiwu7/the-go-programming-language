package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	buffer := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buffer)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}
	return int(*w), nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	buffer := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buffer)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*l++
	}
	return int(*l), nil
}

type Counter struct {
	w     io.Writer
	count int64
}

func (c *Counter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	if err != nil {
		return 0, err
	}
	c.count += int64(n)
	return n, nil
}

// CountingWriter wraps the original io.Writerm and reports
// the number of bytes writter the new Writer at any moment.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wo := &Counter{w: w}
	return wo, &wo.count
}

func main() {
	s := "hello world! \n 你好 世界!"
	var w WordCounter
	var l LineCounter

	fmt.Println("word counter & line counter")
	fmt.Fprint(&w, s) //4
	fmt.Fprint(&l, s) //2

	fmt.Println(w)
	fmt.Println(l)

	fmt.Println("counting writer")
	cw, c := CountingWriter(os.Stdout)
	fmt.Fprintln(cw, "hello")
	fmt.Println("count", *c) // 6
	fmt.Fprintln(cw, "world")
	fmt.Println("count", *c) // 12
}
