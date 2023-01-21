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
	request := Request{
		ID:   123,
		Name: "Зубенко Михаил",
		Cars: []Car{},
		Params: map[string]Param{
			"occupation": {
				ValueID:   850,
				ValueName: "mafia",
			},
		},
	}

	rawJSON, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(rawJSON))
}
