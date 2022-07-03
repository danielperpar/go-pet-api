package application

import(
	"net/http"
	"encoding/json"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"github.com/danielperpar/go-pet-api/infrastructure"
)

var db *sql.DB

type HealthController struct {

}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (controller *HealthController) HealthCheck(writer http.ResponseWriter, request *http.Request) {
	connString := infrastructure.GetConnectionString()

	var errOpen error
	db, errOpen = sql.Open("postgres", connString)
	defer db.Close()

	if errOpen != nil {
		log.Println(errOpen)
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode("unhealthy: BD failed to open")
		return
	}
	
	errPing := db.Ping()
	if errPing != nil {
		log.Println(errPing)
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode("unhealthy: BD failed to ping")
		return
	}
	
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode("healthy")
}