package infrastructure

import ("github.com/danielperpar/go-pet-api/domain")

type IPetRepository interface {
	CreatePet(pet domain.Pet) domain.Pet
	GetPets() []domain.Pet
	GetPredominantSpecies() []string
}