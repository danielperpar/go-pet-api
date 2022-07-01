package application

import (
	"encoding/json"
	"net/http"
)

type PetController struct {
	petService *PetService
}

func NewPetController(petService *PetService) *PetController {
	return &PetController{petService: petService}
}

func (petcontroller *PetController) CreaMascota(writer http.ResponseWriter, request *http.Request) {
	petDto := PetDto{}
	err := json.NewDecoder(request.Body).Decode(&petDto)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	petcontroller.petService.CreatePet(petDto)
	writer.WriteHeader(http.StatusOK) //revisar si mejor un content created con header Location
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(petDto)
}

func (petcontroller *PetController) LisMascotas(writer http.ResponseWriter, request *http.Request) {
	pets := petcontroller.petService.GetPets()
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(pets)
}