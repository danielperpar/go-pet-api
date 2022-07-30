package application

type KpiDto struct {
	PredomSpec *[]string
	AvgAge     float32
	StdDev     float32
}

func NewKpiDto(predomSpec *[]string, avgAge float32, stdDev float32) *KpiDto {
	return &KpiDto{PredomSpec: predomSpec, AvgAge: avgAge, StdDev: stdDev}
}