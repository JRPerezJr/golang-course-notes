package main

import (
	"fmt"
)

type Item struct {
	values []int
}

func (s *Item) Get(idx int) (int, error) {
	if idx > len(s.values) {
		return 0, fmt.Errorf("no element at index %v", idx)
	} else {
		return s.values[idx], nil
	}
}

func main() {
	stuff := Item{}
	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}
}
