package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func count(animal string, n int) {
	defer wg.Done() // decrement counter when done
	for i := 1; i <= n; i++ {
		time.Sleep(time.Duration(400+rand.Intn(200)) * time.Millisecond)
		fmt.Println(i, animal)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(3) // initialize counter to 3
	go count("sheep", 10)
	go count("cat", 10)
	go count("dog", 10)
	wg.Wait() // block until counter is 0

	fmt.Println("zzzZZZzzz")
}
