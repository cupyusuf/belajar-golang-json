package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Produc struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Produc{
		Id:       "1",
		Name:     "Apple",
		ImageURL: "https://www.google.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"1","name":"Apple","image_url":"https://www.google.com"}`
	jsonBytes := []byte(jsonString)

	produc := &Produc{}

	json.Unmarshal(jsonBytes, produc)

	fmt.Println(produc)
	fmt.Println(produc.Id)
	fmt.Println(produc.Name)
	fmt.Println(produc.ImageURL)
}
