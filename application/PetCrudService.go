package application

import (
	"github.com/danielperpar/go-pet-api/domain"
)

type PetCrudService struct {
	petRepository domain.IPetRepository
}

func NewPetCrudService(petRepository domain.IPetRepository) *PetCrudService {
	return &PetCrudService{petRepository: petRepository}
}

func (service *PetCrudService) CreatePet(pet domain.Pet) (domain.Pet, error){
	return service.petRepository.CreatePet(pet)
}

func (service *PetCrudService) GetPets(start, count int) (*[]domain.Pet, error) {
	return service.petRepository.GetPets(start,count)
}