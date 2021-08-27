package main

import (
	"fmt"
	"math/rand"
	"time"
)

func count(animal string, n int) {
	for i := 1; i <= n; i++ {
		time.Sleep(time.Duration(400+rand.Intn(200)) * time.Millisecond)
		fmt.Println(i, animal)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go count("sheep", 10)
	go count("cat", 10)
	go count("dog", 10)

	time.Sleep(5 * time.Second)

	fmt.Println("zzzZZZzzz")
}
