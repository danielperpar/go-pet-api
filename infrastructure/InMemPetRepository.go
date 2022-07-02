package infrastructure

import (
	"github.com/danielperpar/go-pet-api/common"
	"github.com/danielperpar/go-pet-api/domain"
)

var pets []domain.Pet = []domain.Pet{}

type InMemRepository struct{
}

func NewInMemRepository() *InMemRepository {
	return &InMemRepository{}
}

func (repository *InMemRepository) GetPets() *[]domain.Pet{
	return &pets
}

func (repository *InMemRepository) CreatePet(pet domain.Pet) domain.Pet{
	pets = append(pets, pet)
	return pet
}

func (repository *InMemRepository) GetPredominantSpecies() (*[] string, *common.Error) {
	if len(pets) == 0 {
		return nil, &common.Error{Message:common.NO_PETS}
	}

	maxAmount := 0
	speciesAmount := *repository.getSpeciesAmount(&maxAmount)
	
	var species []string;

	for sp,amount := range speciesAmount {
		if amount == maxAmount {
			species = append(species, sp)
		}
	}
	return &species, nil
}

func (repository *InMemRepository) GetAvgAge(species string) (domain.AvgAge, *common.Error) {
	if len(pets) == 0 {
		return domain.AvgAge{}, &common.Error{Message: common.NO_PETS}
	}

	sum := float32(0.0)
	count := float32(0.0)

	for _,pet := range pets {
		if pet.Species == species {
			sum += float32(pet.Age)
			count++
		}
	}
	
	avg := domain.AvgAge{Species: species, Avg: sum/count}
	return avg, nil
}

func (repository *InMemRepository) getSpeciesAmount(outMaxAmount *int) *map[string]int {
	speciesAmount := make(map[string]int)

	count := 0

	for _,pet := range pets {
		_, found := speciesAmount[pet.Species]

		if !found{
			speciesAmount[pet.Species] = 1
		}else{
			speciesAmount[pet.Species] += 1
		}

		if speciesAmount[pet.Species] > count{
			count = speciesAmount[pet.Species]
		}
	}

	*outMaxAmount = count
	return &speciesAmount
}