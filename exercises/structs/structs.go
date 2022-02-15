//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type CalculateArea func(int, int) int
type Area struct {
	length, width int

	// function as field
	calculate CalculateArea
}

type CalculatePerimeter func(int, int, int, int) int

type Perimeter struct {
	side1, side2, side3, side4 int

	getPerimeter CalculatePerimeter
}

type Coordinate struct {
	x, y int
}
type Rectangle struct {
	a Coordinate // top left
	b Coordinate // bottom right
}

func width(rect Rectangle) int {
	return (rect.b.x - rect.a.x)
}
func length(rect Rectangle) int {
	return (rect.a.y - rect.b.y)
}

func area(rect Rectangle) int {
	return length(rect) * width(rect)
}
func perimeter(rect Rectangle) int {
	return (width(rect)*2 + (length(rect) * 2))
}

func printInfo(rect Rectangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}

func main() {
	rect := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printInfo(rect)

	// double the size of the existing reactangle
	rect.a.y *= 2
	rect.b.x *= 2
	printInfo(rect)

	// my solution
	reactangle := Area{
		length: 3,
		width:  3,
		calculate: func(len, width int) int {
			return len * width
		},
	}
	fmt.Println("Rectangle Area")
	fmt.Println("Length", reactangle.length)
	fmt.Println("Width", reactangle.width)
	fmt.Println("Area", reactangle.calculate(reactangle.length, reactangle.width))

	reactangle1 := Perimeter{
		side1: 4,
		side2: 4,
		side3: 8,
		side4: 8,
		getPerimeter: func(side1, side2, side3, side4 int) int {
			return side1 + side2 + side3 + side4
		},
	}
	fmt.Println("Rectangle Perimeter")
	fmt.Println("Side1", reactangle1.side1)
	fmt.Println("Side2", reactangle1.side2)
	fmt.Println("Side3", reactangle1.side3)
	fmt.Println("Side4", reactangle1.side4)
	fmt.Println("Perimeter", reactangle1.getPerimeter(reactangle1.side1, reactangle1.side2, reactangle1.side3, reactangle1.side4))
}
