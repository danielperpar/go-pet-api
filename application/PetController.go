package application

import (
	"net/http"
)

type PetController struct {
	petService *PetService
}

func NewPetController(petService *PetService) *PetController {
	return &PetController{petService: petService}
}

func (petcontroller *PetController) CreaMascota(writer http.ResponseWriter, request *http.Request) {
	//saco el param de la request
	
}

func (petcontroller *PetController) LisMascotas(writer http.ResponseWriter, request *http.Request) {
	//saco el param de la request
	
}