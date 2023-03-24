package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	p := fmt.Println
	p("Packages Log")

	log.Println("test")

	// log.Fatal("Fatal error")

	// log.Panic("Panic error")

	log.Print("Error")

	date := time.Now()

	// file, err := os.Create("myLog.log")
	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))

	if err != nil {
		log.Panic(err)
	}

	l := log.New(file, "Logger: ", log.LstdFlags|log.Llongfile)
	l.Println("pruebas0")
	l.Println("pruebas1")
	l.Println("pruebas2")
	l.Println("pruebas3")

	l.SetPrefix("error: ")
	l.Println("pruebas0")
	l.Println("pruebas1")
	l.Println("pruebas2")
	l.Println("pruebas3")

}
