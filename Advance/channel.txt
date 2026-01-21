// package main

// import (
// 	"fmt"
// 	"time"
// )

// func channel() {
// 	//variable := make([]type, size)
// 	greeting := make(chan string) // create a channel

// 	greetingsting := "Hello, World!"

// 	go func() { // start a goroutine to send data to the channel
// 		greeting <- greetingsting // send a value to the channel
// 		greeting <- "this is golang channel example"
// 		for _, e := range "Golang" {
// 			greeting <- ("alphabeted are " + string(e))
// 		}

// 	}()

// 	reciver := <-greeting // receive a value from the channel
// 	fmt.Println(reciver)
// 	reciver = <-greeting // receive a value from the channel
// 	fmt.Println(reciver)

// 	for range 6 {
// 		reci := <-greeting
// 		fmt.Println(reci)
// 	}

// 	time.Sleep(1 * time.Microsecond)
// 	//fmt.Println(reciver)
// }

package main

import (
	"fmt"
	//"time"
)

func main() {
	chann := make(chan string) // create a channel

	chann <- "Hello World" // send a value to the channel

	go func() {
		reciver := <-chann // receive a value from the channel
		fmt.Println(reciver)
	}()
}
