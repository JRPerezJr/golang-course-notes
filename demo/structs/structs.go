package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	jim := Passenger{"Jim", 1, false}
	fmt.Println(jim)

	var (
		kim    = Passenger{Name: "Kim", TicketNumber: 2}
		sophie = Passenger{Name: "Sophie", TicketNumber: 3}
	)
	fmt.Println(kim, sophie)

	var ada Passenger
	ada.Name = "Ada"
	ada.TicketNumber = 4
	fmt.Println(ada)

	jim.Boarded = true
	kim.Boarded = true
	if jim.Boarded {
		fmt.Println("Jim has boarded the bus")
	}
	if kim.Boarded {
		fmt.Println(kim.Name, "has boarded the bus")
	}

	ada.Boarded = true
	bus := Bus{ada}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
