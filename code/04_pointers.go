package main

import "fmt"

func initPrimes(p *[4]int) {
	// p is a pointer to an array of 4 ints
	p[0] = 2
	p[1] = 3
	p[2] = 5
	p[3] = 7
	fmt.Println(*p) // [2 3 5 7]
}

func main() {
	var x [4]int
	initPrimes(&x)  // &x is a pointer to x
	fmt.Println(x)  // [2 3 5 7]
}
