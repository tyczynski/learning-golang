package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("sadge")
	case "Friday":
		fmt.Println("Bajlando")
	default:
		fmt.Println("Day")
	}

	// alternative to if...else

	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Dzień dobry")
	case hour < 17:
		fmt.Println("Dobrego popołudnia Panu życzę")
	default:
		fmt.Println("Dobry wieczór")
	}

	number := 1
	switch number {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Heloł")
	default:
		fmt.Println("Pyk")
	}
}
