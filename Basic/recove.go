package main

import "fmt"

func main() {

	recoverExample()
	fmt.Println("Finished RecoverExample")

}
func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}

	}() // defer func to recover from panic

	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("After panic") // This line will not be executed
}
