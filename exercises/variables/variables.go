//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var color = "red"
	year, age := 94, 25

	var (
		first = "J"
		last  = "B"
	)
	var days int
	days = (age * 365)
	fmt.Println("Lucky color is", color)
	fmt.Println("Born in", year, "Died at", age)
	fmt.Println("Initials=", first, last)
	fmt.Println("I am", days, "days old")

}
