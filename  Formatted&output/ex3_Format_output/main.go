package main

import(
	"fmt"
)

func main() {
	var ticket_name = "Topgun" //detectเเล้วว่าใช่ type ไหนในการเก็บค่า
	var ticket_byuser = 49
	const ticket_total = 50
	name1 := "polawich"
	fmt.Printf("movie name %v vate 5 stars\n",ticket_name)
	fmt.Printf("Buy 1 ticket amount %v of %v\n",ticket_byuser,ticket_total)
	fmt.Printf("owner name %v\n",name1)
}