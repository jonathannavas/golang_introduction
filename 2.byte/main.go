package main

import (
	"fmt"
)

func main() {
	var A byte = 'A'
	fmt.Println(A)

	var a byte = 'a'
	fmt.Println(a)

	var R byte = 82
	fmt.Println(string(R))

	vector := []byte{65, 97, 82, 115}
	fmt.Println(vector)
	fmt.Println(string(vector))

}
