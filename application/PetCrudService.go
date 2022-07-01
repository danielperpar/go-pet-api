package application

import (
	"github.com/danielperpar/go-pet-api/infrastructure"
)

type PetCrudService struct {
	petRepository infrastructure.IPetRepository
	petMapper *PetMapper
}

func NewPetCrudService(petRepository infrastructure.IPetRepository, petMapper *PetMapper) *PetCrudService {
	return &PetCrudService{petRepository: petRepository, petMapper: petMapper}
}

func (service *PetCrudService) CreatePet(pet PetDto){
	petEntiy := service.petMapper.MapDown(pet)
	service.petRepository.CreatePet(petEntiy)
}

func (service *PetCrudService) GetPets() []PetDto {
	petEntities := service.petRepository.GetPets()
	var petDtos []PetDto = []PetDto{}
	for _,pet:= range petEntities {
		petDto:= service.petMapper.MapUp(pet)
		petDtos = append(petDtos, petDto)
	}
	return petDtos
}

	
		
