package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

func main() {
	http.HandleFunc("/", root)
	fmt.Println("waiting for requests...")
	http.ListenAndServe(":8080", nil)
}

var counter int32

func root(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s\n\n", request)

	count := atomic.AddInt32(&counter, 1)

	fmt.Fprintf(writer, "<html><body><pre>")
	fmt.Fprintln(writer, request.URL.Path)
	fmt.Fprintf(writer, "%d visits\n", count)
	fmt.Fprintf(writer, "</pre></body></html>")
}
