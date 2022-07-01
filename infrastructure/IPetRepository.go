package infrastructure

import ("github.com/danielperpar/go-pet-api/domain")

type IPetRepository interface {
	Create(pet domain.Pet) domain.Pet
	GetPets() []domain.Pet
}