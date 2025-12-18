package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("hello world")
	// to create a map
	myMapp := make(map[string]int) // map[keytype]valuetype
	fmt.Println(myMapp)
	myMapp["key1"] = 10
	myMapp["key2"] = 20
	fmt.Println(myMapp)
	delete(myMapp, "key1")
	fmt.Println(myMapp)

	// to get value from map
	value := myMapp["key2"]
	fmt.Println(value)

	//_, exists := myMapp["key1"] // to check if key exists
	_, unknownvalue := myMapp["key2"] // to check if key exists
	fmt.Println(unknownvalue)         // true

	myMap2 := map[string]int{"a": 1, "b": 1, "c": 3} //map literal
	// map[keytype]valuetype{key1: value1, key2: value2}
	fmt.Println(myMap2)

	//slices := myMap2 ["a":"c"]
	//fmt.Println(slices)

	myMap3 := map[string]int{"a": 1, "b": 1, "c": 3}
	myMap4 := map[int]int{1: 10, 2: 20, 3: 30}
	myMapp5 := map[string]int{"a": 1, "b": 1, "c": 3}
	myMap6 := map[int]int{1: 10, 2: 20, 3: 30}

	if maps.Equal(myMap3, myMapp5) {
		fmt.Println("maps are equal")
	} else {
		fmt.Println("maps are not equal")
	}

	if maps.Equal(myMap4, myMap6) {
		fmt.Println("maps are equal")
	} else {
		fmt.Println("maps are not equal")
	}
	//usig the function i,v for key and value and range is for looping through the map
	for i, r := range myMap3 {
		fmt.Println(i, r)
	}
}
