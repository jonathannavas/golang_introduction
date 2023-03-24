package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	p("Context in golang")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "api-key", "jonathan")
	viewContext(ctx)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go myProcess(ctx)
	select {
	case <-ctx.Done():
		p("deadline")
		p(ctx.Err().Error())
	}

}

func viewContext(ctx context.Context) {
	fmt.Printf("My value is: %s", ctx.Value("api-key"))
}

func myProcess(ctx context.Context) {
	for {
		select {

		case <-ctx.Done():
			fmt.Println("time, out")
			return
		default:
			fmt.Println("doing some process")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
