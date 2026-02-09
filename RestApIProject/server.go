package main

import (
	"fmt"
	"net/http"
	//"strings"
)

func main() {

	port := 3030
	fmt.Printf("Starting server on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
