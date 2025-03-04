package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, err := os.Create("CustomerOut.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)
	customer := Customer{
		FirstName: "Yusuf",
		LastName:  "Supriadi",
		Email:     "yusufsupriadi@gmail.com",
	}

	fmt.Println("Data Berhasil di encode")

	err = encoder.Encode(customer)
	if err != nil {
		panic(err)
	}
}
