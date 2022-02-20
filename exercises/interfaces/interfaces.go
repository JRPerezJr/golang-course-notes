//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftDirector interface {
	DirectVehicle() Lift
}

type Motorcycle string
type Car string
type Truck string
type ModelName string

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Car %v", string(c))

}

func (t Truck) String() string {
	return fmt.Sprintf("Truck %v", string(t))
}

func (m Motorcycle) DirectVehicle() Lift {
	return SmallLift
}

func (c Car) DirectVehicle() Lift {
	return StandardLift
}

func (t Truck) DirectVehicle() Lift {
	return LargeLift
}

func directVehicles(v LiftDirector) {
	switch v.DirectVehicle() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", v)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", v)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", v)
	}
}

func main() {
	car := Car("Ferrari 460")
	truck := Truck("Dodge Ram")
	motorcycle := Motorcycle("BMW K 1600")

	directVehicles(car)
	directVehicles(truck)
	directVehicles(motorcycle)
}
