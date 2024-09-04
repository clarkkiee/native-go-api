package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	db, err := sql.Open("postgres", os.Getenv("DATABASE_CONNECTION_URL"))
	if err != nil {
		panic(err)
	}
	
	log.Println("Connected to Database")
	DB = db

}