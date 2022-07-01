package domain

import "time"

type Pet struct{
	Id	int
	Name string
	Species string
	Gender string
	Age int
	DOB time.Time
}