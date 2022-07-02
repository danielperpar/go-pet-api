package application

import (
	"github.com/danielperpar/go-pet-api/domain"
	"github.com/danielperpar/go-pet-api/infrastructure"
	"github.com/danielperpar/go-pet-api/common"
)

type PetStatsService struct{
	petRepository infrastructure.IPetRepository
}

func NewStatisticsService(petRepository infrastructure.IPetRepository) *PetStatsService{
	return &PetStatsService{petRepository: petRepository}
}

func (service *PetStatsService) GetKpi(species string) (domain.Kpi, *common.Error){
	predSpec,err := service.getPredominantSpecies()
	if err != nil {
		return domain.Kpi{}, &common.Error{Message: err.Message}
	}

	//avgAge,err := service.getAvgAgePerSpecies(species)
	//stdDev,err := service.getStandDevPerSpecies(species)

	kpi := domain.Kpi{PredomSpec: predSpec, AvgAge: domain.AvgAge{}, StdDev: domain.StdDev{}}
	return kpi, nil
}

func (service *PetStatsService)getPredominantSpecies() (*[]string, *common.Error){
	return service.petRepository.GetPredominantSpecies()
}

// func (service *PetStatsService)getAvgAgePerSpecies(species string) (domain.AvgAge, common.Error){
	
// }

// func (service *PetStatsService)getStandDevPerSpecies(species string) (domain.StdDev, common.Error){
	
// }

