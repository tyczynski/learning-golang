package main

import "fmt"

func main() {

	age := 69

	if age == 69 {
		fmt.Println("( ͡° ͜ʖ ͡°)")
	}

	if age > 16 {
		fmt.Println("You're very old man")
	} else {
		fmt.Println("skibidi")
	}

	// initialize and check
	if score := 85; score >= 70 {
		fmt.Println("wooo")
	} else {
		fmt.Println("Still wooo")
	}

	// score is available in if...else scope
	// fmt.Printf("Score: %n\n", score)
}
