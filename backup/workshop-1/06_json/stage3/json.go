package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Request []Command

type Command struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type SendMessageData struct {
	User int64  `json:"user"`
	Text string `json:"text"`
}

type MakeOrderData struct {
	Sku    int64 `json:"sku"`
	Amount int   `json:"amount"`
}

func main() {
	data := `
[
	{
		"type": "send_message",
		"data": {
			"user": 61254895,
			"text": "Hello!"
		}
	},
	{
		"type": "make_order",
		"data": {
			"sku": 12345678,
			"amount": 2
		}
	}
]
`

	var request Request
	err := json.Unmarshal([]byte(data), &request)
	if err != nil {
		log.Fatal(err)
	}

	for _, command := range request {
		switch command.Type {
		case "send_message":
			var data SendMessageData
			err := json.Unmarshal(command.Data, &data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("send_message: %#v\n", data)
		case "make_order":
			var data MakeOrderData
			err := json.Unmarshal(command.Data, &data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("make_order: %#v\n", data)
		}
	}
}
