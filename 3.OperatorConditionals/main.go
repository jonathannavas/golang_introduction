package main

import "fmt"

func main() {

	fmt.Println("Operadores condicionales")

	var isValidAge bool = 5 > 4
	fmt.Println("IsValidAdge:", isValidAge)

	var yearsOld int = 27
	fmt.Println("Can drive:", yearsOld > 18)
	fmt.Println("Can drive:", yearsOld >= 18)
	fmt.Println("Can drive:", yearsOld <= 18)
	fmt.Println("Can drive:", yearsOld < 18)
	fmt.Println("Can drive:", yearsOld == 18)

	fmt.Println()

	//or
	fmt.Println("Can drive:", (yearsOld == 18) || (yearsOld >= 10))

	//and
	fmt.Println("Can drive:", (yearsOld == 18) && (yearsOld >= 10))

	fmt.Println("Not isValidAge:", !isValidAge)

	//condicionales

	yearsOld = 27
	if yearsOld > 18 {
		fmt.Printf("%d Es mayor de 18 \n", yearsOld)
	}

	//another type of declaration variable
	if value := true; value == true {
		fmt.Println("is true value")
	}

	license := true
	age := 15

	if license && age == 19 {
		fmt.Println("Puede seguir avanzando")
	} else if !license && age == 19 {
		fmt.Println("No puede seguir circulando")
	} else if license && age == 15 {
		fmt.Println("No puede seguir circulando")
	}

}
