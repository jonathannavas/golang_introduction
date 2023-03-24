package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println
	p("Paralelismo")

	//para usar paralelismo en varios hilos debemos agregar la palabra reservada llamada go
	go myProcess("A")
	go myProcess("B")

	time.Sleep(3 * time.Second)

	// canales
	myFirstChanel := make(chan string)
	go myProcessWithChanel("C", myFirstChanel)

	result := <-myFirstChanel

	p(result)
	close(myFirstChanel)

	channelA := make(chan string)
	channelB := make(chan string)
	go myProcessWithChanel("D", channelA)
	go myOtherProcessWithChanel("E", channelB)

	resultA := <-channelA
	p(resultA)

	resultB := <-channelB
	p(resultB)

	close(channelA)
	close(channelB)

}

func myProcessWithChanel(p string, c chan string) {
	i := 0
	for i < 20 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
	c <- "ok"
}

func myOtherProcessWithChanel(p string, c chan string) {
	i := 0
	for i < 10 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
	c <- "ok 2"
}
func myProcess(p string) {
	i := 0
	for i < 15 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
}
