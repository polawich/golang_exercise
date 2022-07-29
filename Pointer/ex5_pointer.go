package main

import (
	"fmt"
)

func main(){
	var fname string
	var lname string
	var age int
	fmt.Printf("input yourname\n")
	fmt.Scan(&fname)
	fmt.Printf("input your lastname\n")
	fmt.Scan(&lname)
	fmt.Printf("input your age\n")
	fmt.Scan(&age)

	fmt.Printf("name: %v\nlastname: %v\nage: %v\n",fname,lname,age) // Print ข้อมูบออกมา
} 