package main

import (
	"fmt"
	"math/rand"
	"time"
)

func count(animal string, n int, messages chan<- string) {
	for i := 1; i <= n; i++ {
		time.Sleep(time.Duration(400+rand.Intn(200)) * time.Millisecond)
		messages <- fmt.Sprintf("%d %s", i, animal)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	messages := make(chan string)

	go count("sheep", 10, messages)
	go count("cat", 10, messages)
	go count("dog", 10, messages)

	var deadline <-chan time.Time = time.After(10 * time.Second)
receiving:
	for {
		// blocks until any case is ready (in absence of default case)
		select {
		case message := <-messages:
			fmt.Println(message)

		case <-deadline:
			fmt.Println("deadline")
			break receiving

		case <-time.After(time.Second):
			fmt.Println("channel dead for 1 second")
			break receiving
		}
	}

	fmt.Println("zzzZZZzzz")
}
