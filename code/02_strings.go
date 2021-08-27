package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	s := "Käsebrötchen"
	fmt.Println(s)
	fmt.Println(len(s))              // 14
	fmt.Println(hex.Dump([]byte(s))) // |K..sebr..tchen|

	// loop over the bytes (UTF-8 code units)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%02d: %c\n", i, s[i])
	}
	fmt.Println()
	// loop over the runes (Unicode code points)
	for i, r := range s {
		fmt.Printf("%02d: %c\n", i, r)
	}
}
