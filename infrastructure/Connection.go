package infrastructure

import (	
	"os"
	"database/sql"
	"net/http"

	"github.com/danielperpar/go-pet-api/common"
)

func OpenConnection() (*sql.DB, error){
	connString := GetConnectionString()
	db, errOpen := sql.Open("postgres", connString)
	if(errOpen != nil){
		return nil, common.NewError(http.StatusInternalServerError, errOpen.Error())
	}
	return db, nil
}

func GetConnectionString() string{
	connString := "host=localhost port=5432 user=postgres password=admin dbname=pet_api sslmode=disable"
	if os.Getenv("ENV") == "PROD" {
		connString = os.Getenv("DATABASE_URL")
	}
	return connString
}