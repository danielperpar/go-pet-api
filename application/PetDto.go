package application

import (
	"strings"
	"time"
)

type PetDto struct{
	Id	int 		`json:"-"`
	Name string		`json:"name" example:"luke"`
	Species string	`json:"species" example:"dog"`
	Gender string	`json:"gender" example:"male"`
	Age int			`json:"age" example:"2"`
	Dob time.Time	`json:"dob" example:"2020-05-01T00:00:00Z"`
}

func NewPetDto(name string, species string, gender string, age int, dob time.Time) *PetDto {
	return &PetDto{Name: name, Species: species, Gender: gender, Age: age, Dob: dob}
}

func (p *PetDto) ToLowerCase(pet PetDto) *PetDto{
	return &PetDto{Name : strings.ToLower(pet.Name), 
		Species: strings.ToLower(pet.Species),
		Gender: strings.ToLower(pet.Gender),
		Age: pet.Age,
		Dob: pet.Dob,
	}
}