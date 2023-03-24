package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	p := fmt.Println

	p("Generics")

	v1 := []float64{1.3, 1.4, 6.4, 5.2, 1.58}
	v2 := []int32{1, 4, 32, 56, 7}

	p(Sum(v1))
	p(Sum(v2))

	p(Sum02(v1))
	p(Sum02(v2))

	anyType(1.6, 3.1)
	anyType(1, 3)
	anyType("hola", "mundo")
	anyTypeTwo(1, "mundo")

	comparableType(2, 2)
	comparableType("hola", "hola ")

	orderedValues(2, 3)

	csInt := CustomSlice[int]{1, 2, 3, 4}
	fmt.Println(csInt)
	csString := CustomSlice[string]{"a", "b"}
	fmt.Println(csString)

}

func Sum[N int | int16 | int32 | int64 | float32 | float64](v []N) N {

	var total N
	for _, tv := range v {
		total += tv
	}
	return total

}

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Sum02[N Number](v []N) N {
	var total N
	for _, tv := range v {
		total += tv
	}
	return total
}

func anyType[N any](v1, v2 N) {
	fmt.Println(v1, v2)
}

func anyTypeTwo[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println(v1, v2)
}

// el genérico de tipo comparable nos permite comprobar valores que se pasan por argumentos si ambos son iguales o distintos

func comparableType[N comparable](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 == v2)
}

func orderedValues[N constraints.Ordered](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 == v2)
	fmt.Println(v1 <= v2)
	fmt.Println(v1 >= v2)
}

type CustomSlice[V int | string] []V
