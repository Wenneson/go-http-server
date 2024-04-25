package main

import (
	"log"
	"net/http"
)

func main() {
	// Define the folder to serve the static files from
	fs := http.FileServer(http.Dir("./public"))

	// Handle all requests by serving a file of the same name
	http.Handle("/", fs)

	// Define the address to listen on, use "" to listen on all IP addresses of the machine
	addr := ":3000" // Listening on all interfaces at port 3000

	// Print the address and start the server
	log.Printf("Listening on http://%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
