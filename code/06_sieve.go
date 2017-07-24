package main

import "fmt"

const N = 100

func main() {
	var compound [N]bool
	for i := 2; i*i < N; i++ {
		if !compound[i] {
			for j := i * i; j < N; j += i {
				compound[j] = true
			}
		}
	}
	var primes []int
	for i := 2; i < N; i++ {
		if !compound[i] {
			primes = append(primes, i)
		}
	}
	fmt.Println(primes)
}
