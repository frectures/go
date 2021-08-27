package main

import "fmt"

func initPrimes(a [4]int) {
	// a is an array of 4 ints
	a[0] = 2
	a[1] = 3
	a[2] = 5
	a[3] = 7
	fmt.Println(a)  // [2 3 5 7]
}

func main() {
	var x [4]int    // arrays have a fixed size
	initPrimes(x)   // arrays are passed by value!
	fmt.Println(x)  // [0 0 0 0]
}
