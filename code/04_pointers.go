package main

import "fmt"

func initSha1(a *[20]byte) {
	// a is a pointer to an array of 20 bytes
	a[0] = 0x01
	a[1] = 0x23
	a[2] = 0x45
	a[3] = 0x67
	// ...
	fmt.Println(*a)
}

func main() {
	var x [20]byte
	initSha1(&x) // &x is a pointer to x
	fmt.Println(x)
}
