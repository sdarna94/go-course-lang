package main

import "fmt"

func main() {
	var a, b int = 10, 20

	var result int
	result = a + b

	fmt.Println("result", result)

	result = a - b
	fmt.Println("result", result)

	result = a * b
	fmt.Println("result", result)

	result = a / b
	fmt.Println("result", result)

	result = a % b
	fmt.Println("result", result)

}
