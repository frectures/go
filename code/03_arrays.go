package main

import "fmt"

func initSha1(a [20]byte) {
	// a is an array of 20 bytes
	a[0] = 0x01
	a[1] = 0x23
	a[2] = 0x45
	a[3] = 0x67
	// ...
	fmt.Println(a)
}

func main() {
	var x [20]byte // arrays have a fixed size
	initSha1(x)    // arrays are passed by value!
	fmt.Println(x) // x remains unchanged
}
