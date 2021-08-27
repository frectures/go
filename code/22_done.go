package main

import (
	"fmt"
	"math/rand"
	"time"
)

func count(animal string, n int, messages chan<- string, done chan<- time.Time) {
	defer func() {
		done <- time.Now()
	}()
	for i := 1; i <= n; i++ {
		time.Sleep(time.Duration(400+rand.Intn(200)) * time.Millisecond)
		messages <- fmt.Sprintf("%d %s", i, animal)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	messages := make(chan string)
	done := make(chan time.Time)
	go func() {
		<-done
		<-done
		<-done
		close(messages)
	}()

	go count("sheep", 10, messages, done)
	go count("cat", 10, messages, done)
	go count("dog", 10, messages, done)

	for message := range messages {
		fmt.Println(message)
	}

	fmt.Println("zzzZZZzzz")
}
