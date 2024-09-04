package productcontroller

import (
	"encoding/json"
	"go-crud/config"
	"go-crud/entities"
	"go-crud/utils"
	"net/http"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request) {

	var list_of_products []entities.ProductSchema

	rows, err := config.DB.Query(`SELECT * FROM products`)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var product entities.ProductSchema
		if err := rows.Scan(&product.Id, &product.Product_name, &product.Created_at, &product.Updated_at, &product.Category_id); err != nil {
			panic(err)
		}
		list_of_products = append(list_of_products, product)
	}

	json.NewEncoder(w).Encode(utils.BuildResponseSuccess("Succesfully Fetch all products data", list_of_products))

}