package application

import (
	"encoding/json"
	"net/http"
	"strings"

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

// CreatePet godoc
// @Summary      Create a pet
// @Description  Create a pet providing a pet model
// @Tags         PetController
// @Accept       json
// @Produce      json
// @Param        pet   body      domain.Pet  true  "pet"
// @Success      200  {object}  domain.Pet
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /creamascota [post]
func (petcontroller *PetController) CreatePet(writer http.ResponseWriter, request *http.Request) {
	pet := domain.Pet{}
	err := json.NewDecoder(request.Body).Decode(&pet)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(err.Error())
		return
	}
	pet = pet.ToLowerCase(pet) 
	errCrud := petcontroller.petCrudService.CreatePet(&pet)
	if errCrud != nil{
		custErr := errCrud.(*common.Error)
		writer.WriteHeader(custErr.Code)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(custErr.Message)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(pet)
}

// ListPets godoc
// @Summary      List all pets
// @Description  List all pets in the storage
// @Tags         PetController
// @Accept       json
// @Produce      json
// @Success      200  {array}  domain.Pet
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /lismascotas [get]
func (petcontroller *PetController) ListPets(writer http.ResponseWriter, request *http.Request) {
	pets, errCrud := petcontroller.petCrudService.GetPets()
	if errCrud != nil{
		custErr := errCrud.(*common.Error)
		writer.WriteHeader(custErr.Code)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(custErr.Message)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(*pets)
}

// PetsKpi godoc
// @Summary      Get pets Kpi
// @Description  Get predominant species in the storage. Get average age and std deviation for the provided "species" parameter
// @Tags         PetController
// @Accept       json
// @Produce      json
// @Param        species   query      string  true "species to get average age and std deviation from"
// @Success      200  {object}  domain.Kpi
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /kpidemascotas [get]
func (petcontroller *PetController) PetsKpi(writer http.ResponseWriter, request *http.Request) {
	keys, ok := request.URL.Query()["species"]

	if !ok || len(keys[0]) < 1 {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode("Wrong query parameter")
		return
	}
	petSpecies := keys[0]
	kpi, err := petcontroller.petStatsService.GetKpi(strings.ToLower(petSpecies))
	if err != nil{
		custErr := err.(*common.Error)
		writer.WriteHeader(custErr.Code)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(custErr.Message)
		return
	}

	if err == nil && kpi.PredomSpec == nil {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode("There are no species in the storage yet")
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(*kpi)
}

