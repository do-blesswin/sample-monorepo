package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting webhook server")
	http.HandleFunc("/", WebhookServer)
	http.ListenAndServe(":8080", nil)
}

// WebhookServer responds to webhooks.
func WebhookServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request on webhook server", r.URL.Path)
	fmt.Fprintf(w, "Ack %s", r.URL.Path[1:])
}
