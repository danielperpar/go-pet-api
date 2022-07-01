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
	mapper := application.NewPetMapper()
	service := application.NewPetService(repository, mapper)
	controller := application.NewPetController(service)
	router := mux.NewRouter()
	router.HandleFunc("/creamascota", controller.CreaMascota)
	router.HandleFunc("/lismascotas", controller.LisMascotas)
	router.HandleFunc("/ping", PingHandler)
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