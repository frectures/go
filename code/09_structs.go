package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	myBoss := &Person{"Guido", 60}

	yourBoss := myBoss
	yourBoss.Age++

	fmt.Println(myBoss, yourBoss) // &{Guido 61} &{Guido 61}
}
