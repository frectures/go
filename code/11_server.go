package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	fmt.Println("waiting for requests...")
	http.ListenAndServe(":8080", nil)
}

var counter = 0

func root(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s\n\n", request)
	counter++
	fmt.Fprintf(writer, "<html><body><pre>")
	fmt.Fprintln(writer, request.URL.Path)
	fmt.Fprintf(writer, "%d visits\n", counter)
	fmt.Fprintf(writer, "</pre></body></html>")
}
