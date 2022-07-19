package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	PopCount(1234567890)
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// ex2.3 Rewrite PopCount to use a loop instead of a single
// expression. Compare the performance of the 2 versions.
func PopCountLoop(x uint64) int {
	res := 0
	for i := 0; i <= 7; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

// ex2.4 Write a version of PopCount that ccounts bits by
// shifting its argument through 64 bit positions, testing
// the rightmost bit each time.
func PopCountShift(x uint64) int {
	res := 0
	for i := x; i > 0; i >>= 1 {
		res += int(i & 1)
	}
	return res
}
