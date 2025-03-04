package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "Yusuf",
		LastName:  "Supriadi",
		Email:     "yusuf@gmail.com",
		Age:       25,
		Married:   false,
		Hobbies: []string{
			"Coding",
			"Reading",
			"Gaming",
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{
	"FirstName":"Yusuf",
	"LastName":"Supriadi",
	"Email":"yusuf@gmail.com",
	"Age":25,
	"Married":false,
	"Hobbies":["Coding","Reading","Gaming"]
	}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Email)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	Customer := []Customer{
		{
			FirstName: "Yusuf",
			LastName:  "Supriadi",
			Email:     "yusuf@gmail.com",
			Age:       25,
			Married:   false,
			Hobbies: []string{
				"Coding",
				"Reading",
				"Gaming",
			},
			Addresses: []Address{
				{
					Street:     "Jl. Sudirman",
					Country:    "Indonesia",
					PostalCode: "12345",
				},
				{
					Street:     "Jl. Budhi",
					Country:    "Indonesia",
					PostalCode: "64321",
				},
			},
		},
	}

	jsonBytes, err := json.Marshal(Customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Yusuf","LastName":"Supriadi","Email":"yusuf@gmail.com","Age":25,"Married":false,"Hobbies":["Coding","Reading","Gaming"],"Addresses":[{"Street":"Jl. Sudirman","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jl. Budhi","Country":"Indonesia","PostalCode":"64321"}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Email)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}
