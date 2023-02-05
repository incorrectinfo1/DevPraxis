package main

import (
	"fmt"
)

func askOrder() string {
	var input string
	fmt.Print("What would you like to eat: ")
	// Get the input from the user
	fmt.Scan(&input)
	return input
}

func contains(menu []string, order string) bool {
	for _, item := range menu {
		if item == order {
			return true
		}
	}
	return false
}

func main() {
	fastfoodMenu := []string{"Burgers", "Nuggets", "Fries"}
	fmt.Println("The fast food menu has these items:", fastfoodMenu)

	var total int
	var order string
	// Infinite loop asking for orders
	for order != "quit" {
		order = askOrder()
		if contains(fastfoodMenu, order) {
			total += 2
		} else {
			fmt.Println("Sorry, that item is not on the menu.")
		}
	}

	// Counting $2 bills
	fmt.Println("The total for the order is", total)
}
