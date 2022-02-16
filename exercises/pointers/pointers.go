//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

//* Create functions to activate and deactivate security tags using pointers
func activateTag(tag *SecurityTag) {
	*tag = Active
}
func deactivateTag(tag *SecurityTag) {
	*tag = Inactive
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	fmt.Println("checking out...")
	for i := 0; i < len(items); i++ {
		deactivateTag(&items[i].tag)
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	shirt := Item{"Shirt", Active}
	skirt := Item{"Skirt", Active}
	pants := Item{"Pants", Active}
	jeans := Item{"Jeans", Active}
	//  - Store them in a slice or array
	basket := []Item{shirt, skirt, pants, jeans}
	fmt.Println("Initial", basket)
	//  - Deactivate any one security tag in the array/slice
	deactivateTag(&basket[1].tag)
	fmt.Println("Item 1 deactivated", basket)
	//  - Call the checkout() function to deactivate all tags
	checkout(basket)
	fmt.Println("checked out", basket)
	//  - Print out the array/slice after each change
	activateTag(&basket[1].tag)
	fmt.Println("Item 1 returned", basket)

}
