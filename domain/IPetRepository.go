package domain

type IPetRepository interface {
	CreatePet(pet Pet) Pet
	GetPets() *[]Pet
}