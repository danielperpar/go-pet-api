package infrastructure

import (
	"github.com/danielperpar/go-pet-api/domain"
)

var pets []domain.Pet = []domain.Pet{}

type InMemRepository struct{
}

func NewInMemRepository() *InMemRepository {
	return &InMemRepository{}
}

func (repository *InMemRepository) GetPets() (*[]domain.Pet, error){
	return &pets, nil
}

func (repository *InMemRepository) CreatePet(pet domain.Pet) error{
	pets = append(pets, pet)
	return nil
}
