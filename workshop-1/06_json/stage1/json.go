package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Request struct {
	ID     int              `json:"id"`
	Name   string           `json:"name"`
	Cars   []Car            `json:"cars"`
	Params map[string]Param `json:"params"`
}

type Car struct {
	Plate string `json:"plate"`
	Brand string `json:"brand"`
}

type Param struct {
	ValueID   int64  `json:"value_id"`
	ValueName string `json:"value_name"`
}

func main() {
	data := `
{
	"id": 123,
	"name": "Александр Пушкин",
	"cars": [
		{ "plate":"e123kx777", "brand": "Acura" },
		{ "plate":"a456mp777", "brand": "Ford"  }
	],
	"params": {
		"occupation": {"value_id":57, "value_name":"writer"         },
		"city":       {"value_id":5,  "value_name":"Санкт-Петербург"}
	}
}
`

	var request Request
	err := json.Unmarshal([]byte(data), &request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", request)
}
