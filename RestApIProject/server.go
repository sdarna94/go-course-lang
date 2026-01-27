package main

import (
	"fmt"
	"net/http"
	//"strings"
)

func main() {

	port := ":3030"
	fmt.Printf("Starting server on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}