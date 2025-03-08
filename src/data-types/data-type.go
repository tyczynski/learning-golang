package main

import "fmt"

func main() {
	var integer int = 42
	var floatNum float32 = 3.14
	var greeting string = "Siema Eniu"
	var isPies bool = true

	// type conversion
	var intToFloat float32 = float32(integer)

	fmt.Printf("Integer: %d\n", integer)
	fmt.Printf("Float: %f\n", floatNum)
	fmt.Printf("String: %s\n", greeting)
	fmt.Printf("Boolean: %t\n", isPies)
	fmt.Printf("Converted int to float: %f\n", intToFloat)

	// type information
	fmt.Printf("Type of integer: %T\n", integer)
	fmt.Printf("Type of floatNum: %T\n", floatNum)
}
