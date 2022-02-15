package main

import "fmt"

func main() {
	userList := make(map[string]int)
	userList["Admin"] = 1
	userList["Manager"] = 2
	userList["EndUser"] = 5

	userList["Admin"] += 9
	fmt.Println(userList)

	delete(userList, "Manager")
	fmt.Println("Manager deleted, new user list:", userList)

	fmt.Println("Check user id", userList["Admin"])

	manager, found := userList["Manager"]
	fmt.Println("Check User Manager...")
	if !found {
		fmt.Println("No user present")
	} else {
		fmt.Println("Yes", manager, "present")
	}

	totalIds := 0
	for _, amount := range userList {
		totalIds += amount
	}
	fmt.Println("Total id count", totalIds, "in the DB")
}
