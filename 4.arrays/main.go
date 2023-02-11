package main

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Println("Arrays")
	var myAge int
	fmt.Println(myAge)
	fmt.Printf("Type: %T, bytes: %d, bits: %d\n", myAge, unsafe.Sizeof(myAge), unsafe.Sizeof(myAge)*8)

	var arrAges [6]int
	fmt.Println("Arr ages", arrAges)

	arrNames := [3]string{"jonathan", "gabriel", "navas"}
	fmt.Println("Arr names: ", arrNames)
	fmt.Println("Arr names: ", arrNames[0])

	arrAges[0] = 1
	arrAges[1] = 23
	arrAges[2] = -1523
	arrAges[3] = -1
	arrAges[4] = -4
	fmt.Println("Arr ages", arrAges)

	fmt.Println("Arr names length: ", len(arrNames))

	//vectores == tamaño ya fijado
	//slice == tamaño variable en tiempo de ejecución

	// slices
	var notesUser []int
	fmt.Printf("Notes user: %v, length: %d\n", notesUser, len(notesUser))

	//use append to add more values on array
	notesUser = append(notesUser, 1, 123, 234234, 2, 0, 32)
	fmt.Printf("Notes user updated: %v, length: %d\n", notesUser, len(notesUser))

	mySlice := arrAges[1:4]
	fmt.Println("Copia arr: ", mySlice)
	fmt.Println("Copia arr len: ", len(mySlice))
	fmt.Println(&arrAges[1])
	fmt.Println(&mySlice[0])

	mySlice[0] = 666
	fmt.Println("Copia arr: ", mySlice)
	fmt.Println(&arrAges[1])
	fmt.Println(&mySlice[0])

	slice := make([]int, 3)
	fmt.Println(slice)

}
