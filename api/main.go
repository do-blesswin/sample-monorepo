package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting server")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer says hello.
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request", r.URL.Path)
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
