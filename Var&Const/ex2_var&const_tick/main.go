package main

import(
	"fmt"
)

func main() {
	var ticket_name = "Topgun"
	var ticket_byuser = 49
	const ticket_total = 50
	fmt.Println("movie name:",ticket_name , "rate 5 stars")
	fmt.Println("Buy 1 ticket amount ", ticket_byuser ,"of", ticket_total)
}