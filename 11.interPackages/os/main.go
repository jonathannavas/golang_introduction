package main

import (
	"fmt"
	"os"
)

func main() {

	p := fmt.Println

	p("Package OS")

	file, err := os.Open("myFile.txt")
	if err != nil {
		p(err)
		os.Exit(1)
	}

	data := make([]byte, 200)

	c, err := file.Read(data)

	if err != nil {
		p(err)
		os.Exit(1)
	}

	fmt.Printf("read %d bytes: %q \n ", c, data[:c])

	env1 := os.Getenv("USERNAME")
	p("\nVariable de entorno: ", env1)

	dir, _ := os.Getwd()
	p(dir)

}
