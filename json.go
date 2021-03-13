package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Fruit struct {
	Fruit string `json:"fruit"`
	Size  string `json:"size"`
	Color string `json:"color"`
}

func main() {
	myJson := `
		[
			{
				"fruit": "Apple",
				"size": "Large",
				"color": "Red"
			}, 
			{
				"fruit": "Apple",
				"size": "Large",
				"color": "Red"
			}
		]
	`

	var unmarshalled []Fruit

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Println(unmarshalled)

	// write JSON from a struct
	var mySlice []Fruit

	var m1 Fruit

	m1.Fruit = "barberry"
	m1.Size = "small"
	m1.Color = "red"

	mySlice = append(mySlice, m1)

	var m2 Fruit

	m2.Fruit = "pomgeranate"
	m2.Size = "medium"
	m2.Color = "red"

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", " ")

	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Print(string(newJson))
}
