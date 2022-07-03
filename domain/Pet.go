package domain

import (
	"strings"
	"time"
)

type Pet struct{
	Id	int `json:"-"`
	Name string
	Species string
	Gender string
	Age int
	Dob time.Time
}

func (p Pet) ToLowerCase(pet Pet) Pet{
	return Pet{Name : strings.ToLower(pet.Name), 
		Species: strings.ToLower(pet.Species),
		Gender: strings.ToLower(pet.Gender),
		Age: pet.Age,
		Dob: pet.Dob,
	}
}