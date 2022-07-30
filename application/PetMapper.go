package application

import (
	"github.com/danielperpar/go-pet-api/domain"
)

type PetMapper struct{}

func (mapper *PetMapper) MapDown(pet *PetDto) *domain.Pet {
	return domain.NewPet(pet.Name, pet.Species, pet.Gender, pet.Age, pet.Dob)
}

func (mapper *PetMapper) MapUp(pet *domain.Pet) *PetDto {
	return NewPetDto(pet.Name, pet.Species, pet.Gender, pet.Age, pet.Dob)
}

func (mapper *PetMapper) MapListUp(pets *[]domain.Pet) *[]PetDto{
	petDtos := []PetDto{}

	for _, pet := range *pets {
		petDtos = append(petDtos, *NewPetDto(pet.Name, pet.Species, pet.Gender, pet.Age, pet.Dob))
	}
	return &petDtos
}