package application

import (
	"github.com/danielperpar/go-pet-api/infrastructure"
)

type PetService struct {
	petRepository infrastructure.IPetRepository
	petMapper *PetMapper
}

func NewPetService(petRepository infrastructure.IPetRepository, petMapper *PetMapper) *PetService {
	return &PetService{petRepository: petRepository, petMapper: petMapper}
}

func (service *PetService) Create(pet PetDto){
	petEntiy := service.petMapper.MapDown(pet)
	service.petRepository.Create(petEntiy)
}

func (service *PetService) GetPets() []PetDto {
	petEntities := service.petRepository.GetPets()
	var petDtos []PetDto = []PetDto{}
	for _,pet:= range petEntities {
		petDto:= service.petMapper.MapUp(pet)
		petDtos = append(petDtos, petDto)
	}
	return petDtos
}

	
		
