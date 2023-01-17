package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Tells the net.http package to handle all requests to the web root."/" for giving routes
	http.HandleFunc("/", SimpleServer)
	//
	http.ListenAndServe(":8080", nil)
}

// ResponseWriter sends data to he http client
func SimpleServer(w http.ResponseWriter, r *http.Request) {
	//
	fmt.Fprintf(w, "Hello World %s", r.URL.Path[1:])

}
