package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
	Married   bool
	Hobbies   []string
	Addresses []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Yusuf",
		LastName:  "Supriadi",
		Email:     "yusuf@gmail.com",
		Age:       25,
		Married:   false,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
