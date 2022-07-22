package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h1 := sha256.Sum256([]byte("abc"))
	h2 := sha256.Sum256([]byte("abc"))
	h3 := sha256.Sum256([]byte("a"))
	fmt.Println("Count Bits Demo")
	fmt.Println(compareHashes(h1, h2))
	fmt.Println(compareHashes(h2, h3))
}

// ex4.1
// compareHashes counts the number of bits that are different
// in two SHA256's hash
func compareHashes(h1 [32]byte, h2 [32]byte) int {
	var count int
	for i := 0; i < len(h1); i++ {
		count += countBits(h1[i], h2[i])
	}
	return count
}

// countBits counts the number of bits that are different between
// 2 byte numbers
func countBits(b1 byte, b2 byte) int {
	var count int
	for i := 0; i < 8; i++ {
		bit1 := (b1 >> i) & 1
		bit2 := (b2 >> i) & 1
		if bit1 != bit2 {
			count++
		}
	}
	return count
}
