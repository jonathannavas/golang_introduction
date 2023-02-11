package main

import (
	"fmt"
	"strconv"
	"unsafe"
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

	// constant variables
	const IVA int = 14
	fmt.Println("Iva:", IVA)

	const url string = "https://www.google.com.ec"
	fmt.Println("Url Path: ", url)

	const isValidNumber bool = true
	fmt.Println("IsValidNumber:", isValidNumber)

	//fmt.Printf allow to print with a format
	var my8BitsVariable int8
	fmt.Printf("Int value default %d\n", my8BitsVariable)

	//unsafe package

	fmt.Printf("Int8Bits value default %d\n", my8BitsVariable)
	fmt.Printf("type %T, bytes: %d, bits: %d\n", my8BitsVariable, unsafe.Sizeof(my8BitsVariable), unsafe.Sizeof(my8BitsVariable)*8)

	var myFloat32Value float32
	fmt.Printf("myFloat32Value value default %f\n", myFloat32Value)
	fmt.Printf("type %T, bytes: %d, bits: %d\n", myFloat32Value, unsafe.Sizeof(myFloat32Value), unsafe.Sizeof(myFloat32Value)*8)

	var myFloat64Value float64
	fmt.Printf("myFloat64Value value default %f\n", myFloat64Value)
	fmt.Printf("type %T, bytes: %d, bits: %d\n", myFloat64Value, unsafe.Sizeof(myFloat64Value), unsafe.Sizeof(myFloat64Value)*8)

	fmt.Println()

	var myStringValue string
	fmt.Printf("myStringValue %s\n", myStringValue)
	myStringValue = `hello world`
	fmt.Printf("String value: %s\n", myStringValue)

	myLastName := `
	navas
	cifuentes
	hello
	world`
	fmt.Printf("lastname value: %s\n", myLastName)

	// manage scope
	{
		fmt.Println()
		floatVar := 33.11
		fmt.Printf("type %T, value: %.2f\n", floatVar, floatVar)

		floatStrVar := fmt.Sprintf("%f", floatVar)
		fmt.Printf("type %T, value: %s\n", floatStrVar, floatStrVar)

		intValue1, err := strconv.ParseInt("1223", 0, 64)

		fmt.Println()
		fmt.Println(err)
		fmt.Printf("type %T, value: %d\n", intValue1, intValue1)

		intValue2, err := strconv.ParseInt("1223.321", 0, 64)

		fmt.Println()
		fmt.Println(err)
		fmt.Printf("type %T, value: %d\n", intValue2, intValue2)

		floatValue1, _ := strconv.ParseFloat("1223.123123123123", 64)

		fmt.Println()
		fmt.Println(err)
		fmt.Printf("type %T, value: %f\n", floatValue1, floatValue1)
	}

}
