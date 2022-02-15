//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(title string, servers map[string]int) {
	fmt.Println()
	fmt.Println("---", title, "---")
	//* Create a function to print server status, including:
	//  - Number of servers
	//  - Number of servers for each status (Online, Offline, Maintenance, Retired)

	// My solution
	// numOfOnline := 0
	// numOfOffline := 0
	// numOfUnderMaintenace := 0
	// numOfRetired := 0
	// for _, server := range servers {

	// 	if server == 0 {
	// 		numOfOnline += 1
	// 	} else if server == 1 {
	// 		numOfOffline += 1
	// 	} else if server == 2 {
	// 		numOfUnderMaintenace += 1
	// 	} else {
	// 		numOfRetired += 1
	// 	}
	// }

	// fmt.Println("Number of servers", len(servers), "\nOnline", numOfOnline, "\nOffline", numOfOffline, "\nUnder maintenance", numOfUnderMaintenace, "\nRetired", numOfRetired)

	// Cleaner solution
	fmt.Println("Number of servers", len(servers))

	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}
	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are undergoing maintenance")
	fmt.Println(stats[Retired], "servers are retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverList := make(map[string]int)
	for _, server := range servers {
		serverList[server] = Online
	}
	//* Perform the following status changes and display server info:
	//  - display server info
	printServerStatus("First operation", serverList)

	//  - change `darkstar` to `Retired`
	//  - change `aiur` to `Offline`
	//  - display server info
	serverList["darkstar"] = Retired
	serverList["aiur"] = Offline
	printServerStatus("Second operation", serverList)

	//  - change all servers to `Maintenance`
	//  - display server info
	for _, server := range servers {
		serverList[server] = Maintenance
	}
	printServerStatus("Third operation", serverList)

}
