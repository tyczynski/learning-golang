package main

import "fmt"

func main() {
	// go has only `for` loop

	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration: %d\n", i)
	}

	// for as while
	whileCounter := 0
	for whileCounter < 3 {
		fmt.Println("Counter:", whileCounter)
		whileCounter++
	}

	// infinite loop
	sum := 0
	for {
		sum++
		if sum > 10 {
			break
		}
	}
	fmt.Println("Sum:", sum)

	// for-range (iter over collections)
	fruits := []string{"apple", "banana", "orange"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}
}
