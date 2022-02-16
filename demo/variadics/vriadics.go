package main

import "fmt"

func sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	all := append(a, b...)

	answer := sum(all...)
	fmt.Println("Answer 1:", answer)

	answer2 := sum(1, 2, 3, 4, 5, 6)
	fmt.Println("Answer 2:", answer2)
}
