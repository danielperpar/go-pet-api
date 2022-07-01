package application

import (
	"fmt"

	"github.com/danielperpar/go-pet-api/domain"
	"github.com/danielperpar/go-pet-api/infrastructure"
)

type PetStatsService struct{
	petRepository infrastructure.IPetRepository
	petMapper *PetMapper
}

func NewStatisticsService(petRepository infrastructure.IPetRepository, petMapper *PetMapper) *PetStatsService{
	return &PetStatsService{petRepository: petRepository, petMapper: petMapper}
}

func (service *PetStatsService) GetKpi(species string) domain.Kpi{
	predSpec := service.getPredominantSpecies()
	//avgAgePerSpec := service.getAvgAgePerSpecies(species)
	//standDevSpec := service.getStandDevPerSpecies(species)

	fmt.Println(species)

	kpi := domain.Kpi{
		PredomSpec: predSpec,
		AvgAgePerSpec: domain.AvgAgePerSpecies{"dummy",1},
		StandDevPerSpec: domain.StandDevPerSpecies{"dummy",1},
	}

	fmt.Println(kpi)

	return kpi
}

func (service *PetStatsService)getPredominantSpecies() string{
	pets := service.petRepository.GetPets()

	speciesAmount := make(map[string]int)

	for _,pet := range pets {
		_, found := speciesAmount[pet.Species]

		if(!found){
			speciesAmount[pet.Species] = 1
		}else{
			speciesAmount[pet.Species] += 1
		}
	}
	
	max := 0
	species := ""

	for sp,amount := range speciesAmount {
		if amount > max {
			max = amount
			species = sp	
		}
	}
	
	return species
 }

func (service *PetStatsService)getAvgAgePerSpecies(species string) *domain.AvgAgePerSpecies {
	return &domain.AvgAgePerSpecies{Species: "dummy", Avg: 0}
}

func (service *PetStatsService)getStandDevPerSpecies(species string) *domain.StandDevPerSpecies {
	return &domain.StandDevPerSpecies{Species: "dummy", StandDev: 0}
}

func findPet(pets []domain.Pet, id int) *domain.Pet{
	for _,pet:= range pets {
		if(pet.Id == id){
			return &pet
		}
	}
	return nil
}

