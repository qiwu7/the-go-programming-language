package main

import (
	"57/the-go-programming-language/ch3/complex/mandelbrot"
	"fmt"
)

func main() {
	fmt.Println("Complex Number Demo")
	complexDemo()

	fmt.Println("Mandelbrot Demo")
	mandelbrot.Run()
}

func complexDemo() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Printf("x = %v, y + %v\n", x, y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("real(x * y) = ", real(x*y))
	fmt.Println("imag(x * y) = ", imag(x*y))

	fmt.Println("1i * 1i = ", 1i*1i)
}
