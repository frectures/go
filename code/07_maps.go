package main

import "fmt"

func main() {
	birth := map[string]int{
		"C":    1972,
		"C++":  1983,
		"Java": 1994,
	}
	for language, year := range birth {
		fmt.Printf("%s was released in %d\n", language, year)
	}
	birth["Go"] = 2007
	examine(birth, "Go")
	delete(birth, "Go")
	examine(birth, "Go")
}

func examine(birth map[string]int, language string) {
	if year, present := birth[language]; present {
		fmt.Printf("%s was released in %d\n", language, year)
	} else {
		fmt.Printf("Never heard of %s...\n", language)
	}
}
