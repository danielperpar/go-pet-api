package application

import (
	"encoding/json"
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

// ShowAccount godoc
// @Summary      Create a pet
// @Description  Create a pet providing a pet model
// @Tags         PetController
// @Accept       json
// @Produce      json
// @Param        pet   body      domain.Pet  true  "pet"
// @Success      200  {object}  domain.Pet
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /creamascota [post]
func (petcontroller *PetController) CreaMascota(writer http.ResponseWriter, request *http.Request) {
	pet := domain.Pet{}
	err := json.NewDecoder(request.Body).Decode(&pet)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(err.Error())
		return
	}
	pet,errCrud := petcontroller.petCrudService.CreatePet(pet)
	if errCrud != nil{
		custErr := errCrud.(*common.Error)
		writer.WriteHeader(custErr.Code)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(custErr.Message)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(pet)
}

func (petcontroller *PetController) LisMascotas(writer http.ResponseWriter, request *http.Request) {
	pets, err := petcontroller.petCrudService.GetPets()
	if err != nil{
		//manageErrors(err, writer)
	}
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
		//manageErrors(err, writer)
	}
	
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(kpi)
}