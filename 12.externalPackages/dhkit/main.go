package main

import (
	"encoding/json"
	"fmt"

	"github.com/digitalhouse-tech/go-lib-kit/meta"
	"github.com/digitalhouse-tech/go-lib-kit/response"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func main() {
	p := fmt.Println
	p("Package DhKit")

	u := &User{
		FirstName: "Jonathan",
		LastName:  "Navas",
		Email:     "jonathan@google.com",
		Age:       27,
	}

	p(u)

	print(response.OK("", u, nil, nil))
	print(response.OK("Todo bien", u, nil, nil))

	pagination := meta.New(1, 30, 100)
	print(response.OK("Todo bien", u, pagination, nil))

	print(response.Created("", u, nil, nil))
	print(response.Accepted("", u, nil, nil))

	print(response.BadRequest("My error"))
	print(response.BadRequest(""))
	print(response.NotFound("My error"))
	print(response.NotFound(""))
	print(response.InternalServerError("My error"))
	print(response.InternalServerError(""))

}

func print(v interface{}) {
	j, _ := json.Marshal(v)
	fmt.Println(string(j))
	fmt.Println()
}
