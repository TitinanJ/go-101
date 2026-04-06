package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	value := Person{Name: "Alice", Age: 30}
	jsonString, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	
	var alice Person
	err = json.Unmarshal(jsonString, &alice)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonString))
	fmt.Println(alice)
}