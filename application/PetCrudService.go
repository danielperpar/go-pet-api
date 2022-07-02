package application

import (
	"github.com/danielperpar/go-pet-api/infrastructure"
	"github.com/danielperpar/go-pet-api/domain"
)

type PetCrudService struct {
	petRepository infrastructure.IPetRepository
}

func NewPetCrudService(petRepository infrastructure.IPetRepository) *PetCrudService {
	return &PetCrudService{petRepository: petRepository}
}

func (service *PetCrudService) CreatePet(pet domain.Pet) domain.Pet{
	return service.petRepository.CreatePet(pet)
}

func (service *PetCrudService) GetPets() *[]domain.Pet {
	pets := service.petRepository.GetPets()
	return pets
}

	
		
