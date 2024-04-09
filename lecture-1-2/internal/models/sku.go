package models

// SKU - единица складского учёта
type SKU struct {
	ID    SKUID  `json:"item_id" ,db:"id"`     // ID товарной единицы
	Name  string `json:"item_name" ,db:"name"` // Название товарной единицы
	Price uint64 `json:"item_price" ,db:"-"`   // Цена одной товарной единицы
	/* ... */
}
