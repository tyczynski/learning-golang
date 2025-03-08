package main

import "fmt"

func main() {
	// explicit declaration
	var name string = "Przemek"

	// type inference => var age = 28
	age := 28

	// multiple variable declaration
	var (
		city           string = "Koniec Świata"
		styrtaIsFiring bool   = false
		weight                = 130
	)

	// constants
	const PI = 3.14

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Styrta się dopaliła?:", styrtaIsFiring)
	fmt.Println("Weight:", weight)
}
