package infrastructure

import (	
	"os"
	"fmt"
)

func GetConnectionString() string{
	var connString string

	var env string = os.Getenv("ENV")
	if env == ""{
		connString = "host=localhost port=5432 user=postgres password=admin dbname=pet_api sslmode=disable"
	} else {
	user := os.Getenv("USER")
	password := os.Getenv("PASS")
	connString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	"heroku", 5432, user, password, "pet_api")
	}

	return connString
}