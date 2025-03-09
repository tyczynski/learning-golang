package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Printf("Hello %s! ", name)

	switch {
	case age < 13:
		fmt.Println("You're too young.")
	case age >= 13 && age < 60:
		fmt.Println("Here is your ticket.")
	default:
		fmt.Println("You're too old.")
	}
}
