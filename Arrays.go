package main

import "fmt"

func main() {
	//var arrayName [size]dataType
	var numbers [5]int
	fmt.Println(numbers)

	numbers[0] = 10
	fmt.Println(numbers)
	numbers[4] = 20
	fmt.Println(numbers)

	// Declare and initialize an array
	//var arrayName = [size]dataType{value1, value2, ..., valueN}
	frutis := [3]string{"Apple", "Banana", "Mango"} // Array of strings
	fmt.Println("No of frutis ", frutis)            // Output: [Apple Banana Mango]

	oldfrutis := frutis // Copying array
	// Displaying both arrays
	// Modifying oldfrutis will not affect frutis
	oldfrutis[0] = "Orange" // Changing first element of oldfrutis
	// Displaying both arrays
	fmt.Println("No of frutis ", oldfrutis) // Output: [Orange Banana Mango]
	fmt.Println("No of frutis ", frutis)    // Output: [Apple Banana Mango]
}
