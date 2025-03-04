package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirsName string
	LastName string
	Email    string
	Age      int
	Married  bool
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirsName: "Yusuf",
		LastName: "Supriadi",
		Email:    "yusuf@gmail.com",
		Age:      25,
		Married:  false,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
