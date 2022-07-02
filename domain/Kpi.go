package domain

type AvgAgePerSpecies struct {
	Species string
	Avg     int
}

type StandDevPerSpecies struct {
	Species  string
	StandDev float32
}

type Kpi struct {
	PredomSpec      []string
	AvgAgePerSpec   AvgAgePerSpecies
	StandDevPerSpec StandDevPerSpecies
}