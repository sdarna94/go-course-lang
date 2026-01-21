package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 42
	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Println("Original value:", x)
	fmt.Println("Value:", v)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	fmt.Println("Is value valid?", v.Kind() == reflect.Int)

}
