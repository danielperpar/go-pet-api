package application

import (
	"github.com/danielperpar/go-pet-api/domain"
)

type KpiMapper struct {}

func (*KpiMapper) MapUp(kpi *domain.Kpi) *KpiDto {
	return NewKpiDto(kpi.PredomSpec, kpi.AvgAge, kpi.StdDev)
}
