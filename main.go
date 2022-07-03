package main

import (
	"github.com/danielperpar/go-pet-api/application"
	"github.com/danielperpar/go-pet-api/domain"
	"github.com/danielperpar/go-pet-api/infrastructure"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

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
