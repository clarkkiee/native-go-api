package entities

import "time"

type (
	ProductSchema struct {
		Id           int8      `json:"id"`
		Product_name string    `json:"product_name"`
		Created_at   time.Time `json:"created_at"`
		Updated_at   time.Time    `json:"updated_at"`
		Category_id  int8      `json:"category_id"`
	}
)