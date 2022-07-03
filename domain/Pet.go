package domain

import "time"

type Pet struct{
	Id	int `json:"-"`
	Name string
	Species string
	Gender string
	Age int
	Dob time.Time
}