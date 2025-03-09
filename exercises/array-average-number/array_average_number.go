package main

import "fmt"

func main() {
	grades := [5]float64{13, 12.5, 2, 20, 4}
	fmt.Printf("Grades: %v\n", grades)

	var sum float64
	var max float64
	for _, num := range grades {
		sum += num

		if num > max {
			max = num
		}
	}

	average := sum / float64(len(grades))
	fmt.Printf("Average grade: %v\n", average)
	fmt.Printf("The best grade: %f\n", max)
}
