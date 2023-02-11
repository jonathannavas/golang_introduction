package main

import (
	"fmt"
)

func main() {
	fmt.Println("For use")

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 1

	for sum < 1000 {
		sum++
		fmt.Println("Contador: ", sum)
	}

	sum = 0
	for {
		if sum > 1000 {
			break
		}
		sum++
	}
	fmt.Println("Valor de suma: ", sum)

	arrAge := []int{51, 22, 3, 14, 5, 63}

	for i := 0; i < len(arrAge); i++ {
		fmt.Printf("Position: %d, valor: %d \n", i, arrAge[i])
	}

	fmt.Println("Otra forma de usar el ciclo for")

	for i := range arrAge {
		fmt.Printf("Position: %d, valor: %d \n", i, arrAge[i])
	}

	fmt.Println("Otra forma de usar el ciclo for con valor")
	// llave indice valor muy similar al map
	for i, v := range arrAge {
		fmt.Printf("Position: %d, valor: %d \n", i, v)
	}

	// recorrer mapas
	map2 := map[string]string{
		"name":     "jonathan",
		"lastname": "navas",
		"email":    "jonathan@google.com",
	}

	for key, value := range map2 {
		fmt.Printf("llave: %s, valor: %s \n", key, value)
	}

	// recorrer matrices bidimensionales
	map3 := map[string][]int{
		"a": nil,
		"b": {2, 1, 123, 14, 5},
		"c": {4, 5, 6, 7, 8, 9, 10},
	}

	for key, value := range map3 {
		fmt.Println("Key: ", key)
		for _, v := range value {
			fmt.Println("Value: ", v)
		}
		fmt.Println()
	}

	//Dentro de nuestro código ya tenemos definido un array, debemos reccorrerlo e incrementar todos sus valores en 20.

	array := [5]int{4, 2, 5, 6, 7}

	// realizar la funcionalidad

	for i, v := range array {
		array[i] = v + 20
	}

	fmt.Println("Los valores del array son: ", array)

	title := `Ingresa un valor`
	fmt.Println(title)

	var arrNums []int
	var num int
	for {
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		arrNums = append(arrNums, num)
	}
	fmt.Println("Numeros ingresados: ", arrNums)

	fmt.Println("Tienda consulta")

	arrProducts := map[string]string{
		"10": "notebook",
		"15": "tv",
		"21": "heladera",
		"27": "monitor",
		"35": "camara",
	}

	var opt string
	var searchTerms []string
	for {
		fmt.Scanln(&opt)
		if opt == "0" {
			break
		}

		val, ok := arrProducts[opt]
		if !ok {
			searchTerms = append(searchTerms, "No encontrado")
		} else {
			searchTerms = append(searchTerms, val)
		}
	}

	fmt.Println("Tus busquedas fueron: ", searchTerms)

}
