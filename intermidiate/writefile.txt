package main

import (
	"fmt"
	"os"
)

func main() {
	FILE, ERR := os.Create("test.txt")
	if ERR != nil {
		fmt.Println("Error creating file:", ERR)
		return
	}
	defer FILE.Close()
	_, ERR = FILE.WriteString("Hello, World!\n")
	if ERR != nil {
		fmt.Println("Error writing to file:", ERR)
		return
	}
}
