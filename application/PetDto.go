package application

import "time"

type PetDto struct{
	Id	int
	Name string
	Species string
	Gender string
	Age int
	Dob time.Time
}