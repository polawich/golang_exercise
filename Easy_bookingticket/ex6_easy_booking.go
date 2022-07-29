package main

import (
	"fmt"
)

func main() {
	var fname string
	var lname string
	var ticket_buy int
	const ticket_total = 50

	fmt.Println("Welcome to Ticket Movie Booking!")
	fmt.Printf("input your name: ")
	fmt.Scan(&fname)
	fmt.Printf("input your lastname: ")
	fmt.Scan(&lname)
	fmt.Printf("input your ticket buy: ")
	fmt.Scan(&ticket_buy)

	total_fortick := ticket_total - ticket_buy

	fmt.Printf("name: %v\nlastname: %v\nticket: %v\nremaining: %v for %v\n", fname, lname, ticket_buy, total_fortick, ticket_total)
}
