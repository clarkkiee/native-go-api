package entities

import "time"

type (
	CategorySchema struct {
		Id            int8      `json:"id"`
		Created_at    time.Time `json:"created_at"`
		Updated_at    time.Time    `json:"updated_at"`
		Category_name string    `json:"category_name"`
	}
)