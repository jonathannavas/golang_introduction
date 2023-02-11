package main

import (
	"fmt"

	"github.com/jonathannavas/golang_introduction/7.functions/function"
)

func main() {
	fmt.Println("Funciones")

	sum := function.Add(5, 4)
	fmt.Println("Suma de dos números:", sum)
	fmt.Println()
	fmt.Println("Llamado directo a la función")
	fmt.Println("Suma de otros dos números:", function.Add(3, 2))

	function.RepeatString(10, "Hola")

	val, err := function.Calc(function.MUL, 3, 8)
	fmt.Printf("Resultado: %.0f, error: %v\n", val, err)

	val, err = function.Calc(function.DIV, 3, 0)
	fmt.Printf("Resultado: %.0f, error: %v\n", val, err)

	val, err = function.Calc(function.SUB, 3, 1)
	fmt.Printf("Resultado: %.0f, error: %v\n", val, err)

	val1, val2 := function.Split(20)
	fmt.Printf("Value 1: %d, value 2: %d\n", val1, val2)

	valorSum := function.SumData(1, 2, 34, 54, 56, 445, 645, 645, 6)
	fmt.Println("Sum values total: ", valorSum)

	val, err = function.MultipleOperations(function.DIV, 3, 2, 1, 1, 4, 5)
	fmt.Printf("Resultado: %.2f, error: %v\n", val, err)

	fmt.Println("Factory pattern")

	fn := function.FactoryOperation(function.SUM)
	v := fn(1, 3)
	fmt.Println("Factory pattern value sum: ", v)

	fn = function.FactoryOperation(function.DIV)
	v = fn(1, 0)
	fmt.Println("Factory pattern value div: ", v)

}
