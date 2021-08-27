package main

import "fmt"

func initPrimes(s []int) {
	// s is a slice (view) into an array of ints
	s[0] = 2
	s[1] = 3
	s[2] = 5
	s[3] = 7
	fmt.Println(s)  // [2 3 5 7]
}

func main() {
	var x [8]int
	// x[b:e] is a slice into x, starting at b, ending at e
	// b defaults to 0, e defaults to len(x)
	initPrimes(x[2:6])
	fmt.Println(x)  // [0 0 2 3 5 7 0 0]
}
