package domain

type IPetRepository interface {
	CreatePet(pet Pet) (Pet,error)
	GetPets(start, count int) (*[]Pet,error)
}