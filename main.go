package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/danielperpar/go-pet-api/application"
	_ "github.com/danielperpar/go-pet-api/docs"
	"github.com/danielperpar/go-pet-api/domain"
	"github.com/danielperpar/go-pet-api/infrastructure"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Pet API
// @version         1.0
// @description     Create pets and obtain stats from the pet DB
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
	repository := infrastructure.NewPostgrePetRepositoy()
	err := repository.OpenConnection()
	if err != nil{
		log.Fatal(err)
	}
	petCrudService := application.NewPetCrudService(repository)
	petStatsService := domain.NewStatisticsService(repository)
	controller := application.NewPetController(petCrudService, petStatsService)
	healthController := application.NewHealthController()
	router := mux.NewRouter()
	router.HandleFunc("/creamascota", controller.CreaMascota).Methods("POST")
	router.HandleFunc("/lismascotas", controller.LisMascotas).Methods("GET")
	router.HandleFunc("/kpidemascotas", controller.KpiDeMascotas).Methods("GET")
	router.HandleFunc("/health", healthController.HealthCheck).Methods("GET")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("entro en la api")})

	swaggerUrl := "http://localhost:8080/swagger/doc.json"
	port := "8080"
	addr := "127.0.0.1" + ":" + port

	if os.Getenv("ENV") == "PROD" {
		host := "0.0.0.0" 
		port = os.Getenv("PORT")
		addr = host + ":" + port
		swaggerUrl = addr + "/swagger/doc.json"
		
		log.Println("address =>" + addr)
		log.Println("swagger =>" + swaggerUrl)
	}

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(swaggerUrl),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	
	http.Handle("/", router)

	server := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on :" + port + " ...")
	log.Fatal(server.ListenAndServe())
}
