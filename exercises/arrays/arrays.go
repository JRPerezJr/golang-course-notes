//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type ShoppingList struct {
	name  string
	price int
}

func printItems(list [4]ShoppingList) {
	var totalItems, totalPrice int
	for i := 0; i < len(list); i++ {
		item := list[i]
		if item.name != "" {
			totalItems++
		}
		totalPrice += item.price
	}
	fmt.Println("Last item on the list:", list[totalItems-1])
	fmt.Println("Total items:", totalItems)
	fmt.Println("Total price $", totalPrice)
}

func main() {
	items := [4]ShoppingList{
		{"Soap", 3},
		{"Milk", 4},
		{"Meat", 20},
	}
	printItems(items)

	items[3] = ShoppingList{"Beer", 20}
	printItems(items)
}
