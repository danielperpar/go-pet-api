package main

import (
	"github.com/danielperpar/go-pet-api/application"
	"github.com/danielperpar/go-pet-api/infrastructure"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	repository := infrastructure.NewInMemRepository()
	petCrudService := application.NewPetCrudService(repository)
	petStatsService := application.NewStatisticsService(repository)
	controller := application.NewPetController(petCrudService, petStatsService)
	router := mux.NewRouter()
	router.HandleFunc("/creamascota", controller.CreaMascota).Methods("POST")
	router.HandleFunc("/lismascotas", controller.LisMascotas).Methods("GET")
	router.HandleFunc("/kpidemascotas", controller.KpiDeMascotas).Methods("GET")
	router.HandleFunc("/ping", PingHandler).Methods("GET")
	http.Handle("/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on :" + port + " ...")
	log.Fatal(server.ListenAndServe())
}

func PingHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode("pong")
}