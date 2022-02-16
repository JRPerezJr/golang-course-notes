//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

const (
	RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
)

type MemberName string
type BookTitle string

type BookStatus bool

type LendAudit struct {
	checkOut string
	checkIn  string
}

type Member struct {
	name  MemberName
	books map[BookTitle]LendAudit
}
type BookEntry struct {
	totalBooks int
	lentOut    int
}

type Library struct {
	members map[MemberName]Member
	books   map[BookTitle]BookEntry
}

func printTime() string {
	dt := time.Now()
	format := dt.Format(time.RFC1123)
	return format
}

func printBookAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn == "" {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut, "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printBookAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/total", book.totalBooks, "/lent out:", book.lentOut)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title BookTitle, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of the library")
		return false
	}
	if book.lentOut == book.totalBooks {
		fmt.Println("No more books available to lend")
		return false
	}
	book.lentOut += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: printTime()}
	return true
}

func returnBook(library *Library, title BookTitle, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}
	book.lentOut -= 1
	library.books[title] = book

	audit.checkIn = printTime()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books:   make(map[BookTitle]BookEntry),
		members: make(map[MemberName]Member),
	}

	library.books["The Go Lander"] = BookEntry{
		totalBooks: 4,
		lentOut:    0,
	}
	library.books["Little Go Riding Hood"] = BookEntry{
		totalBooks: 3,
		lentOut:    0,
	}
	library.books["Lord of the Go"] = BookEntry{
		totalBooks: 2,
		lentOut:    0,
	}
	library.books["Return of the Go"] = BookEntry{
		totalBooks: 1,
		lentOut:    0,
	}
	library.books["Go strikes back"] = BookEntry{
		totalBooks: 5,
		lentOut:    0,
	}

	library.members["James"] = Member{"James", make(map[BookTitle]LendAudit)}
	library.members["Vesper"] = Member{"Vesper", make(map[BookTitle]LendAudit)}
	library.members["Jaws"] = Member{"Jaws", make(map[BookTitle]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.members["James"]
	checkedOut := checkoutBook(&library, "Return of the Go", &member)
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Return of the Go", &member)
	fmt.Println("\nCheck in a book:")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
