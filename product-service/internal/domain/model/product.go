package model

type Product struct {
	SKU   int64  `json:"sku"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}
