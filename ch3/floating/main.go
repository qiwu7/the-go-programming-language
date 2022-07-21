package main

import (
	"57/the-go-programming-language/ch3/floating/surface"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Exponent Demo")
	exponent()

	fmt.Println("Special Values Demo")
	specialValues()

	fmt.Println("Surface")
	surface.Surface()
}

// %g, %f, %e
func exponent() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d, ex = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func specialValues() {
	var z float64
	fmt.Println("z", z)
	fmt.Println("-z", -z)
	fmt.Println("1/z", 1/z)
	fmt.Println("-1/z", -1/z)
	fmt.Println("z/z", z/z)
}
