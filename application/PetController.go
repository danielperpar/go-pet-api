package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/danielperpar/go-pet-api/common"
	"github.com/danielperpar/go-pet-api/domain"
)

type PetController struct {
	petCrudService *PetCrudService
	petStatsService *domain.PetStatsService
}

func NewPetController(petCrudService *PetCrudService, petStatsService *domain.PetStatsService) *PetController {
	return &PetController{petCrudService: petCrudService, petStatsService: petStatsService}
}

func (petcontroller *PetController) CreaMascota(writer http.ResponseWriter, request *http.Request) {
	pet := domain.Pet{}
	err := json.NewDecoder(request.Body).Decode(&pet)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	petcontroller.petCrudService.CreatePet(pet)
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(pet)
}

func (petcontroller *PetController) LisMascotas(writer http.ResponseWriter, request *http.Request) {
	pets := petcontroller.petCrudService.GetPets()
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(*pets)
}

func (petcontroller *PetController) KpiDeMascotas(writer http.ResponseWriter, request *http.Request) {
	
	keys, ok := request.URL.Query()["species"]

	if !ok || len(keys[0]) < 1 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	petSpecies := keys[0]
	kpi, err := petcontroller.petStatsService.GetKpi(petSpecies)

	if err != nil{
		switch err.Error() {
		case common.NoPets :
			fmt.Fprintf(writer, "%v", common.NoPets)
			writer.WriteHeader(http.StatusNoContent)
		default:
			fmt.Fprintf(writer, "%v", common.UnknownError)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(kpi)
}
