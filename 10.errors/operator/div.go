package operator

func Div(x, y float64) float64 {
	// defer func() {
	// 	fmt.Println("in my div function defer")
	// }()
	if y <= 0 {
		panic("solo usar numeros positivos y y debe ser mayor a cero")
	}
	return x / y
}
