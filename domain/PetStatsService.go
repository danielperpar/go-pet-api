package domain

import (
	"errors"
	"github.com/danielperpar/go-pet-api/common"
	"math"
)

type PetStatsService struct{
	petRepository IPetRepository
}

func NewStatisticsService(petRepository IPetRepository) *PetStatsService{
	return &PetStatsService{petRepository: petRepository}
}

func (service *PetStatsService) GetKpi(species string) (Kpi, error){
	pets := service.petRepository.GetPets()
	if len(*pets) == 0 {
		return Kpi{}, errors.New(common.NoPets)
	}
	predSpec := service.getPredominantSpecies(pets)
	avgAge := service.getAvgAge(pets, species)
	stdDev := service.getStandDev(pets, avgAge, species)

	kpi := Kpi{PredomSpec: predSpec, AvgAge: avgAge, StdDev: stdDev}
	return kpi, nil
}

func (service *PetStatsService) getPredominantSpecies(pets *[]Pet) *[] string {
	maxAmount := 0
	speciesAmount, maxAmount := service.getSpeciesAmount(pets)
	
	var species []string;

	for sp,amount := range *speciesAmount {
		if amount == maxAmount {
			species = append(species, sp)
		}
	}
	return &species
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

func (service *PetStatsService) getAvgAge(pets *[]Pet, species string) float32 {	
	sum := float32(0.0)
	count := float32(0.0)

	for _,pet := range *pets {
		if pet.Species == species {
			sum += float32(pet.Age)
			count++
		}
	}
	return sum/count
}

 func (service *PetStatsService)getStandDev(pets *[]Pet, avgAge float32, species string) float32 {
	var sum float64 = 0.0

	for _,pet := range *pets {
		if pet.Species == species {
			diff := (float64(avgAge) - float64(pet.Age))
			pow := math.Pow(diff,2) 
			sum += pow
		}
	}
	return float32(math.Sqrt(sum))
 }

