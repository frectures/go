package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func count(animal string, n int, messages chan<- string) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		time.Sleep(time.Duration(400+rand.Intn(200)) * time.Millisecond)
		messages <- fmt.Sprintf("%d %s", i, animal)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	messages := make(chan string)
	wg.Add(3)
	go func() {
		wg.Wait()
		close(messages)
	}()

	go count("sheep", 10, messages)
	go count("cat", 10, messages)
	go count("dog", 10, messages)

	for {
		message, sentBeforeClose := <-messages
		if !sentBeforeClose {
			fmt.Println("channel closed")
			break
		}
		fmt.Println(message)
	}

	fmt.Println("zzzZZZzzz")
}
