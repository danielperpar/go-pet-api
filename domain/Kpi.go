package domain

type Kpi struct {
	PredomSpec *[]string
	AvgAge     float32
	StdDev     float32
}

func NewKpi(predomSpec *[]string, avgAge float32, stdDev float32) *Kpi {
	return &Kpi{PredomSpec: predomSpec, AvgAge: avgAge, StdDev: stdDev} 
}
