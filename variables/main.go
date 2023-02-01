package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world variables")
	var age int
	age = 27
	fmt.Println("My Age", age)

	var positiveValue uint
	positiveValue = 1
	fmt.Println("This variable only works with posittive values", positiveValue)

	var name string
	name = "Jonathan"
	fmt.Println("My name is:", name)

	var isPaid bool
	isPaid = true
	fmt.Println("The order is paid?: ", isPaid)

	fmt.Println("Address memory variable", &isPaid)

	// implicit declaration variable
	lastName := "Navas"
	fmt.Println("My last name:", lastName)

	oldAge := 27
	fmt.Println("Another tye of declaration:", oldAge)

}
