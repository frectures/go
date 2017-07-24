package main

import "fmt"

func initSha1(a []byte) {
	// a is a slice (view) into an array of bytes
	a[0] = 0x01
	a[1] = 0x23
	a[2] = 0x45
	a[3] = 0x67
	// ...
	fmt.Println(a)
}

func main() {
	var x [20]byte
	initSha1(x[:]) // x[:] is a slice (view) into x
	fmt.Println(x) // x[:] is short-hand for x[0:20]
}
