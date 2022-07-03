package domain

type IPetRepository interface {
	CreatePet(pet Pet) error
	GetPets() (*[]Pet,error)
}