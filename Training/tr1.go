package main

import(
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(math.Pi)
	fmt.Println(ret1(1,5))

	a, b := ret2("world", "hello")
	fmt.Println(a, b)

	var c, d = ret3(15)
	fmt.Println(c, d)

	var e, f int = 1, 2
	fmt.Println(e, f)

	var tr, fls = true, false
	fmt.Println(tr, fls)
}

// add
func ret1(a ,b int) int {
	return a+b
}

// swap
func ret2(a, b string) (string, string) {
	return b, a
}

// naked return
func ret3(sum int) (x, y int) {
	x = sum * 2 / 4
	y = sum - x
	// return (x, y)
	return
}


func types() (x bool) {
	/*
		bool string int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr 
		byte // alias for uint8 
		rune // alias for int32 // represents a Unicode code point 
		float32 float64 complex64 complex128
	*/
	x = false
	return 
}