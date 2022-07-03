package infrastructure

import (	
	"os"
)

func GetConnectionString() string{
	connString := "host=localhost port=5432 user=postgres password=admin dbname=pet_api sslmode=disable"
	if os.Getenv("ENV") == "PROD" {
		connString = os.Getenv("DATABASE_URL")
	}
	return connString
}