package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapDecode(t *testing.T) {
	jsonString :=
		`{
		"id":1,
		"name":"Iphone 14",
		"price":20000000
		}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonByte, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestJsonMapEncode(t *testing.T) {
	result := map[string]interface{}{
		"id":    1,
		"name":  "Iphone 14",
		"price": 20000000,
	}

	jsonByte, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonByte))
}
