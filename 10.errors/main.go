package main

import (
	"fmt"

	"errors"

	"github.com/jonathannavas/golang_introduction/10.errors/operator"
)

func main() {
	fmt.Println("Error managment")
	var err error
	err = errors.New("my custom error")
	// si necesito el valor del error en tipo string debo utilizar el metodo Error()
	fmt.Println(err)
	fmt.Println(err.Error())

	err2 := fmt.Errorf("my format err, string: %s, number: %d", "test error", 1)
	fmt.Println(err2)
	fmt.Println(err2.Error())

	defer func() {
		fmt.Println("defer main function")
		r := recover()
		if r != nil {
			fmt.Println("not results")

			fmt.Println("Recovered in: ", r)
		}
	}()

	z := operator.Div(4, -2)

	// el defer siempre se ejecuta al finalizar la funcion
	fmt.Println("results")

	fmt.Println("Division:", z)

}
