package application

import (
	"encoding/json"
	"net/http"
)

type PetController struct {
	petCrudService *PetCrudService
	petStatsService *PetStatsService
}

func NewPetController(petCrudService *PetCrudService, petStatsService *PetStatsService) *PetController {
	return &PetController{petCrudService: petCrudService, petStatsService: petStatsService}
}

func (petcontroller *PetController) CreaMascota(writer http.ResponseWriter, request *http.Request) {
	petDto := PetDto{}
	err := json.NewDecoder(request.Body).Decode(&petDto)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	petcontroller.petCrudService.CreatePet(petDto)
	writer.WriteHeader(http.StatusOK) //revisar si mejor un content created con header Location
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(petDto)
}

func (petcontroller *PetController) LisMascotas(writer http.ResponseWriter, request *http.Request) {
	pets := petcontroller.petCrudService.GetPets()
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(pets)
}

func (petcontroller *PetController) KpiDeMascotas(writer http.ResponseWriter, request *http.Request) {
	
	keys, ok := request.URL.Query()["species"]

	if !ok || len(keys[0]) < 1 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	petSpecies := keys[0]

	kpi := petcontroller.petStatsService.GetKpi(petSpecies)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(kpi)
}
