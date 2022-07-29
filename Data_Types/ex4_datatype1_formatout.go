package main

import (
	"fmt"
)

func main() {
	var user string
	var age int
	user = "polawich"
	age = 12
	fmt.Printf("ชื่อ %v อายุ %v\n",user,age)
	fmt.Printf("ชื่อ %T อายุ %T\n",user,age) // printlog ว่าเป็น Type อะไรด้วย %T
}