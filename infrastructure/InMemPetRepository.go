package infrastructure

import ("github.com/danielperpar/go-pet-api/domain")

var Pets []domain.Pet = []domain.Pet{}

type InMemRepository struct{
}

func (repository InMemRepository) GetPets() []domain.Pet{
	return Pets
}

func (repository InMemRepository) Create(pet domain.Pet){
	Pets = append(Pets, pet)
}

