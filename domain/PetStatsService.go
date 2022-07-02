package domain

import (
	"github.com/danielperpar/go-pet-api/common"
)

type PetStatsService struct{
	petRepository IPetRepository
}

func NewStatisticsService(petRepository IPetRepository) *PetStatsService{
	return &PetStatsService{petRepository: petRepository}
}

func (service *PetStatsService) GetKpi(species string) (Kpi, *common.Error){
	pets := service.petRepository.GetPets()
	if len(*pets) == 0 {
		return Kpi{}, &common.Error{Message:common.NO_PETS}
	}
	
	predSpec,err := service.getPredominantSpecies(pets)
	if err != nil {
		return Kpi{}, &common.Error{Message: err.Message}
	}
	avgAge,err := service.getAvgAge(pets, species)
	if err !=nil {
		return Kpi{}, &common.Error{Message: err.Message}
	}

	//stdDev,err := service.getStandDevPerSpecies(species)

	kpi := Kpi{PredomSpec: predSpec, AvgAge: avgAge, StdDev: StdDev{}}
	return kpi, nil
}

func (service *PetStatsService) getPredominantSpecies(pets *[]Pet) (*[] string, *common.Error) {
	maxAmount := 0
	speciesAmount, maxAmount := service.getSpeciesAmount(pets)
	
	var species []string;

	for sp,amount := range *speciesAmount {
		if amount == maxAmount {
			species = append(species, sp)
		}
	}
	return &species, nil
}

func (service *PetStatsService) getSpeciesAmount( pets *[]Pet) (*map[string]int, int)  {
	speciesAmount := make(map[string]int)

	count := 0

	for _,pet := range *pets {
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

	
	return &speciesAmount,count
}

func (service *PetStatsService) getAvgAge(pets *[]Pet, species string) (AvgAge, *common.Error) {
	sum := float32(0.0)
	count := float32(0.0)

	for _,pet := range *pets {
		if pet.Species == species {
			sum += float32(pet.Age)
			count++
		}
	}
	
	avg := AvgAge{Species: species, Avg: sum/count}
	return avg, nil
}


// func (service *PetStatsService)getStandDevPerSpecies(species string) (domain.StdDev, common.Error){
	
// }

