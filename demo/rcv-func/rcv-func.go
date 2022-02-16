package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot1 := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println("Initial:", lot1)
	lot1.occupySpace(1)
	occupySpace(&lot1, 2) //without receiver function
	fmt.Println("After occupied:", lot1)

	lot1.vacateSpace(2)
	fmt.Println("After vacate:", lot1)
}
