package domain

type IPetRepository interface {
	CreatePet(pet Pet) (Pet,error)
	GetPets() (*[]Pet,error)
}