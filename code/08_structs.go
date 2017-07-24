package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	alice := Person{"Alice", 21}

	bob := alice
	bob.Name = "Bob"
	bob.Age++

	fmt.Printf("%v\n", alice) // {Alice 21}
	fmt.Printf("%#v\n", bob)  // main.Person{Name:"Bob", Age:22}
}
