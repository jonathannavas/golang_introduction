package main

import (
	"fmt"
	"strconv"
)

func main() {

	p := fmt.Println

	p("Package strconv")

	s := strconv.Itoa(120)
	p(s)

	s = strconv.FormatBool(true)
	p(s)

	s = strconv.FormatFloat(3.14, 'E', -1, 64)
	p(s)

	//	el primer numero del método corresponde al número a convertir y el segundo parámetro es la base a la que se piensa modificar en este caso va del
	// 42 a binario
	s = strconv.FormatInt(-42, 2)
	p(s)

	s = strconv.FormatUint(42, 2)
	p(s)

	val, err := strconv.ParseBool("true")
	p("val:", val, "err:", err)

	val1, err := strconv.ParseFloat("3.1416", 64)
	p("val1:", val1, "err:", err)

	val2, err := strconv.ParseInt("-42", 10, 64)
	p("val2:", val2, "err:", err)

	val3, err := strconv.ParseUint("42", 10, 64)
	p("val3:", val3, "err:", err)

}
