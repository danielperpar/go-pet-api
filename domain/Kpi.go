package domain

type AvgAge struct {
	Species string
	Avg     float32
}

type StdDev struct {
	Species  string
	StandDev float32
}

type Kpi struct {
	PredomSpec *[]string
	AvgAge     AvgAge
	StdDev     StdDev
}
