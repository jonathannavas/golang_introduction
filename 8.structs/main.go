package main

import (
	"encoding/json"
	"fmt"

	"github.com/jonathannavas/golang_introduction/8.structs/structsInterface/structs"
	"github.com/jonathannavas/golang_introduction/8.structs/structsInterface/vehicles"
)

func main() {
	fmt.Println("Structs")

	// 1 way to declare a struct

	var p1 structs.Product

	p1.ID = 1
	p1.Name = "TV"

	p1.Price = 314.12

	fmt.Println("Product: ", p1)

	// 2 another way to declare a struct
	p2 := structs.Product{}
	p2.ID = 2
	p2.Name = "Laptop"
	p2.Type = structs.Type{
		ID:          1,
		Code:        "laptop",
		Description: "Descripción de la laptop más moderna",
	}
	p2.Price = 1320.99

	fmt.Println("Product: ", p2)

	// 3 another way to declare a struct

	p3 := structs.Product{3, "Iphone 14", structs.Type{}, 13, 1456.33, nil}
	fmt.Println("Product: ", p3)

	p4 := structs.Product{ID: 4, Name: "Tablet", Price: 423.11}
	fmt.Println("Product: ", p4)

	p5 := structs.Product{
		ID:   5,
		Name: "Congelador Wilson",
		Type: structs.Type{
			ID:          111,
			Code:        "cg1",
			Description: "Electrodoméstico",
		},
		Price: 145.63,
		Tags:  []string{"heladeria", "helados", "postres"},
	}
	fmt.Println("Product: ", p5)

	p6 := structs.Product{
		ID:   6,
		Name: "Monitor Gamer Aorus",
		Type: structs.Type{
			ID:          111,
			Code:        "cg2",
			Description: "monitor",
		},
		Count: 2,
		Price: 145.63,
		Tags:  []string{"monitor", "gamer"},
	}
	fmt.Println("Product: ", p6)

	valor, err := json.Marshal(p6)
	fmt.Println(err)
	fmt.Println(string(valor))
	fmt.Println("Precio del p6: ", p6.GetTotalPrice())
	p6.SetName("Botella")
	fmt.Println("Producto cambiado el nombre: ", p6)
	p6.AddTags("Prueba", "nuevo tag")
	fmt.Println("Producto agregando mas tags: ", p6)

	p7 := structs.Product{
		ID:   6,
		Name: "Monitor Gamer Asus",
		Type: structs.Type{
			ID:          111,
			Code:        "cg2",
			Description: "monitor",
		},
		Count: 1,
		Price: 12.63,
		Tags:  []string{"monitor", "gamer"},
	}

	p8 := structs.Product{
		ID:   6,
		Name: "Monitor Gamer LG",
		Type: structs.Type{
			ID:          111,
			Code:        "cg2",
			Description: "monitor",
		},
		Count: 1,
		Price: 135.3,
		Tags:  []string{"monitor", "gamer"},
	}

	cart := structs.NewCart(1231)
	cart.AddProducs(p6, p7, p8)

	fmt.Println("...Carrito de compras...")
	fmt.Println("Total Products:", len(cart.Products))
	fmt.Printf("Total cart: $%.2f\n", cart.Total())

	fmt.Println()
	fmt.Println("Vehicles")

	carV := vehicles.Car{Time: 120}
	fmt.Println("Distance:", carV.Distance())

	vArray := []string{"car", "moto", "truck", "test"}
	d := 400
	for _, v := range vArray {
		vehi, err := vehicles.New(v, d)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println()
			continue
		}
		fmt.Printf("Vehicle: %s,Distance: %.2f\n", v, vehi.Distance())
	}

}
