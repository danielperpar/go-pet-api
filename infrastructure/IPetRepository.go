package infrastructure

import (
	"github.com/danielperpar/go-pet-api/domain"
	"github.com/danielperpar/go-pet-api/common"
)

type IPetRepository interface {
	CreatePet(pet domain.Pet) domain.Pet
	GetPets() *[]domain.Pet
	GetPredominantSpecies() (*[]string, *common.Error)
	GetAvgAge(species string) (domain.AvgAge, *common.Error)
}