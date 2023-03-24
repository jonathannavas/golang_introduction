package main

import "fmt"

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("Add: %p", &m)
	m.Name = name
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("Add: %p", m)
	m.Name = name
}

func main() {
	fmt.Println("Punteros")

	var age int
	var pointerAge *int
	age = 27
	pointerAge = &age
	fmt.Println("Age:", age)
	fmt.Println("Age reference:", &age)
	fmt.Println("Age pointer reference:", pointerAge)
	fmt.Println("Age pointer:", *pointerAge)

	*pointerAge = 10
	fmt.Println("Age pointer:", *pointerAge)

	fmt.Println()
	fmt.Println("------ANOTHER EXAMPLE------------")

	myVar := 30
	fmt.Printf("Address: %p, values: %v\n", &myVar, myVar)

	increment(myVar)
	fmt.Println(myVar)

	fmt.Println()
	incrementPointer(&myVar)
	fmt.Println(myVar)
	fmt.Println()

	var mySlice []int
	mySlice = append(mySlice, 2, 3, 4)
	fmt.Printf("Address %p, values:%v\n", mySlice, mySlice)
	fmt.Printf("Address1 %p, Address2:%p, Address3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Println(mySlice)
	fmt.Printf("Address1 %p, Address2:%p, Address3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])

	fmt.Println()
	fmt.Println()
	fmt.Println("Struct")

	// opcion 1 para manejo de punteros con estructuras
	myStruct := &MyStruct{ID: 1, Name: "Jonathan"}
	fmt.Printf("addrs %p\n", myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

	updateStruct(myStruct)
	fmt.Println()

	fmt.Printf("addrs %p\n", myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

	// // opcion 2 para manejo de punteros con estructuras
	// myStruct := MyStruct{ID: 1, Name: "Jonathan"}
	// // fmt.Printf("addrs %p\n", myStruct)
	// fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

	// updateStruct(&myStruct)
	// fmt.Println()

	// // fmt.Printf("addrs %p\n", myStruct)
	// fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

	fmt.Println()
	fmt.Printf("address %p \n", myStruct)
	myStruct.Set("navas")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	myStruct.SetP("cifuentes")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

}

func increment(val int) {
	fmt.Println(&val)
	val++
}

func incrementPointer(val *int) {
	fmt.Println(val)
	*val++
}

func updateSlice(v []int) {
	fmt.Printf("Inside Address: %p\n", v)
	fmt.Printf("Inside Address 1: %p, address 2: %p, address 3: %p\n", &v[0], &v[1], &v[2])
	v[0] = 12
}

func updateStruct(v *MyStruct) {
	fmt.Printf("add in function: %p", v)
	v.ID = 101
	v.Name = "Gabriel"
}
