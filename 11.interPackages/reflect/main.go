package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
}

func main() {
	p := fmt.Println

	p("Package reflect")

	// myInt := 5
	// myPnt := &myInt
	usr := User{1, "jonathan@google.com", "Jonathan", "Navas"}
	Examine(usr)

}

func Examine(data interface{}) {

	if reflect.ValueOf(data).Kind() == reflect.Struct {
		v := reflect.ValueOf(data)
		fmt.Print("Number of fields:", v.NumField())
		fmt.Println()
		for i := 0; i < v.NumField(); i++ {

			fmt.Printf("Field:%d type:%s value:%v\n", i, v.Field(i).Kind(), v.Field(i))

			switch v.Field(i).Kind() {
			case reflect.Int, reflect.Int64, reflect.Int32:
				myVar := v.Field(i).Int()
				fmt.Printf("Field:%d type:%T value:%v\n", i, myVar, myVar)
			}
		}
	} else {

		t := reflect.TypeOf(data)
		v := reflect.ValueOf(data)
		k := t.Kind()
		fmt.Println("Type:", t)
		fmt.Println("Value:", v)
		fmt.Println("Kind:", k)

	}

}
