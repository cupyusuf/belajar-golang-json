package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonString := `{
	"FirstName":"Yusuf",
	"LastName":"Supriadi",
	"Email":"yusuf@gmail.com",
	"Age":25,
	"Married":false
	}`

	jsonByte := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Email)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
