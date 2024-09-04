package main

import (
	"go-crud/config"
	"go-crud/controller/healthcheckcontroller"
	"go-crud/controller/productcontroller"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	http.HandleFunc("/status/db", healthcheckcontroller.DatabaseHealthCheck)
	http.HandleFunc("/products", productcontroller.GetAllProduct)

	log.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)

}