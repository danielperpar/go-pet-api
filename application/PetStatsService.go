package application

import (
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

	kpi := domain.Kpi{
		PredomSpec: predSpec,
		AvgAgePerSpec: domain.AvgAgePerSpecies{"dummy",1},
		StandDevPerSpec: domain.StandDevPerSpecies{"dummy",1},
	}

	return kpi
}

func (service *PetStatsService)getPredominantSpecies() []string{
	return service.petRepository.GetPredominantSpecies()
}

func (service *PetStatsService)getAvgAgePerSpecies(species string) *domain.AvgAgePerSpecies {
	return &domain.AvgAgePerSpecies{Species: "dummy", Avg: 0}
}

func (service *PetStatsService)getStandDevPerSpecies(species string) *domain.StandDevPerSpecies {
	return &domain.StandDevPerSpecies{Species: "dummy", StandDev: 0}
}

