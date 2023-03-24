package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	p := fmt.Println
	p("Package uuid")

	id := uuid.New().String()
	p(id)

	id2 := uuid.NewString()
	p(id2)

	uid := uuid.New()
	p(uid.Version())
	p(uid.Variant())

}
