package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// //create a new http client
	// client := &http.Client{}

	// //make a get request
	// resp, err := client.Get("http://jsonplaceholder.typicode.com/posts/1")

	// //handle error
	// if err != nil {
	// 	fmt.Println("error in code ", err)

	// }
	// //close the body at the end of the function
	// defer resp.Body.Close()

	// //read the body
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("error in code ", err)
	// }
	// fmt.Println(body)
	// fmt.Println(string(body))

	client := &http.Client{}
	//resp, err := client.Get("http://jsonplaceholder.typicode.com/posts/1")
	resp, err := client.Get("https://swapi.dev/api/people/?page=2")
	if err != nil {
		fmt.Println("error in code ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in code ", err)
	}
	//fmt.Println(body)
	fmt.Println(string(body))
}
