package application

import (
	"github.com/danielperpar/go-pet-api/domain"
)

type PetCrudService struct {
	petRepository domain.IPetRepository
	petMapper *PetMapper
}

func NewPetCrudService(petRepository domain.IPetRepository) *PetCrudService {
	return &PetCrudService{petRepository: petRepository}
}

func (service *PetCrudService) CreatePet(pet *PetDto) error {
	return service.petRepository.CreatePet(service.petMapper.MapDown(pet))
}

func (service *PetCrudService) GetPets() (*[]PetDto, error) {
	pets, err := service.petRepository.GetPets()
	return service.petMapper.MapListUp(pets), err
}