package main

import "fmt"

func arrays() {
	fmt.Println("--- Arrays ---")

	// fixed length
	var numbers [5]int

	// same type
	colors := [3]string{"Red", "Green", "Blue"}

	// ... counts the elements for me
	animals := [...]string{"Dog", "Cat", "Bird", "Bober"}

	fmt.Printf("Numbers: %d, Colors: %v, Animals: %v", numbers, colors, animals)
}

func slices() {
	fmt.Println("--- Slices ---")

	// with initial values
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Numbers: %v\n", numbers)

	// create a slice with make(type, length, capacity)
	slice := make([]int, 3, 5)
	fmt.Printf("Slice length: %d, capacity: %d\n", len(slice), cap(slice))
}

func maps() {
	fmt.Println("--- Maps ---")

	// empty map
	scores := make(map[string]int)
	fmt.Printf("Scores: %v\n", scores)

	// initialize map
	population := map[string]int{
		"Poznań":   500000,
		"Olsztyn":  1500,
		"Koszalin": 134,
	}
	fmt.Printf("Population: %v\n", population)

	// if value exists
	if pop, exists := population["Poznań"]; exists {
		fmt.Println("Poznań population: ", pop)
	}

	// add or update value
	population["Poznań"] = 12

	// delete a key
	delete(population, "Poznań")

	// iteration
	for city, pop := range population {
		fmt.Printf("%s, %d people\n", city, pop)
	}
}

func structs() {
	fmt.Println("--- Maps ---")

	type Person struct {
		Name    string
		Age     int
		Address string
	}

	grzegorz := Person{
		Name:    "Grzegorz Brzęczyszczykiewicz",
		Age:     45,
		Address: "Chrząszczyżewoszyce, powiat Łękołody",
	}

	fmt.Println(grzegorz.Address)
	fmt.Println(grzegorz.Age)
	fmt.Println(grzegorz.Name)
}

func main() {
	arrays()
	slices()
	maps()
	structs()
}
