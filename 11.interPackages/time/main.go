package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)

	then := time.Date(2023, 3, 9, 20, 34, 21, 42342, time.UTC)
	p(then)

	then = then.Add(time.Hour * 1)
	p(then)

	then = then.Add(time.Second * 12)
	p(then)

	then = then.Add(time.Hour * -1)
	p(then)

	p(then.Weekday())

	diff := now.Sub(then)
	p(diff.Hours())

}
