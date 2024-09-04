package healthcheckcontroller

import (
	"encoding/json"
	"go-crud/config"
	"go-crud/utils"
	"net/http"
)

func DatabaseHealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := config.DB.Query(`SELECT 1`)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(utils.BuildResponseSuccess("Database is up and runnning", nil))
}