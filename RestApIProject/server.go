package main

import (
	"fmt"
	"net/http"
	//"strings"
)

func main() {
	// Define the handler for the /Order endpoint
	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Order endpoint called")
	})

	port := 3030
	fmt.Printf("Starting server on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

}
