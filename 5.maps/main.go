package main

import "fmt"

func main() {
	fmt.Println("Maps")

	m1 := make(map[int]string)
	m1[1] = "a"
	m1[2] = "b"
	m1[3] = "c"

	fmt.Println(m1)
	fmt.Println(m1[1])

	delete(m1, 2)
	fmt.Println(m1)

	m1[1] = "aa"
	m1[100] = "aa"
	fmt.Println(m1)

	// si no existe muestra el valor por defecto de cada tipo
	fmt.Println(m1[123])

	v1, ok := m1[1]
	fmt.Printf("value:%s, ok:%v\n", v1, ok)

	v2, ok := m1[10000]
	fmt.Printf("value:%s, ok:%v\n", v2, ok)

	m2 := map[int]string{1: "a", 2: "h", 3: "d"}
	fmt.Println(m2)

}
