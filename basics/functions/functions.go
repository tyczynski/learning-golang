package main

import (
	// "errors"
	"fmt"
)

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, errors.New("division by zero")
// 	}

// 	return a / b, nil
// }

// func calculateStats(numbers []int) (min, max, sum int) {
// 	if len(numbers) == 0 {
// 		return // returns min=0, max=0, sum=0
// 	}

// 	for _, n := range numbers {
// 		if n < min {
// 			min = n
// 		}
// 		if n > max {
// 			max = n
// 		}
// 		sum += n
// 	}

// 	return // returns current values
// }

// Methods and interfaces

type Car struct {
	Make    string
	Model   string
	Year    int
	Mileage float64
}

type Vehicle interface {
	DisplayInfo() string
	Drive(distance float64)
}

func (c Car) DisplayInfo() string {
	return fmt.Sprintf("%d %s %s with %.1f miles", c.Year, c.Make, c.Model, c.Mileage)
}

func (c *Car) Drive(distance float64) {
	c.Mileage += distance
}

type Motorcycle struct {
	Make    string
	Model   string
	Year    int
	Mileage float64
}

func (c Motorcycle) DisplayInfo() string {
	return fmt.Sprintf("%d %s %s with %.1f miles", c.Year, c.Make, c.Model, c.Mileage)
}

func (c *Motorcycle) Drive(distance float64) {
	c.Mileage += distance
}

func TakeTrip(v Vehicle, distance float64) {
	fmt.Println("Before trip:", v.DisplayInfo())
	v.Drive(distance)
	fmt.Println("After trip:", v.DisplayInfo())
}

func main() {
	// myCar := Car{
	// 	Make:    "Toyota",
	// 	Model:   "Corolla",
	// 	Year:    2020,
	// 	Mileage: 15000,
	// }

	// fmt.Println(myCar.DisplayInfo())

	// myCar.Drive(100)
	// fmt.Println(myCar.DisplayInfo())

	myCar := Car{Make: "Toyota", Model: "Corolla", Year: 2020, Mileage: 15000}
	myMotorcycle := Motorcycle{Make: "Honda", Model: "CBR", Year: 2019, Mileage: 8000}

	// Both Car and Motorcycle can be used as Vehicle
	TakeTrip(&myCar, 100)
	TakeTrip(&myMotorcycle, 100)
}
