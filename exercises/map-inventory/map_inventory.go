package main

import "fmt"

var inventory = make(map[string]int)

func addProduct(name string, amount int) {
	if _, exists := inventory[name]; exists {
		inventory[name] += amount
	} else {
		inventory[name] = amount
	}
}

func removeProduct(name string) {
	if _, exists := inventory[name]; exists {
		delete(inventory, name)
	} else {
		fmt.Println("There is no product named: ", name)
	}
}

func removeProductAmount(name string, amount int) {
	if val, exists := inventory[name]; exists {
		val -= amount

		if val <= 0 {
			removeProduct(name)
		} else {
			inventory[name] = val
		}
	} else {
		fmt.Println("There is no product named: ", name)
	}
}

func displayProducts() {
	fmt.Println("--- Current Inventory ---")
	for name, amount := range inventory {
		fmt.Printf("* %s, amount %d\n", name, amount)
	}
}

func main() {
	addProduct("Jabłko", 15)
	addProduct("Kiełbasa", 4)
	addProduct("Jabłko", 4)
	displayProducts()

	removeProductAmount("Kiełbasa", 2)
	displayProducts()

	removeProductAmount("Kiełbasa", 4)
	displayProducts()

	addProduct("Kiełbasa", 4)
	removeProduct("Kiełbasa")
	displayProducts()
}
