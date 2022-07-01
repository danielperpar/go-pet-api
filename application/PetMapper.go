package application

import ("github.com/danielperpar/go-pet-api/domain")

type PetMapper struct{

}

func (service *PetMapper) MapDown (petDto PetDto) domain.Pet{
	petEntity := domain.Pet{
		Id: petDto.Id,
		Name : petDto.Name, 
		Species: petDto.Species, 
		Gender: petDto.Gender, 
		Age: petDto.Age, 
		Dob: petDto.Dob,
	}
	return petEntity
}

func (service *PetMapper) MapUp (pet domain.Pet) PetDto{
	petDto := PetDto{
		Id: pet.Id,
		Name : pet.Name, 
		Species: pet.Species, 
		Gender: pet.Gender, 
		Age: pet.Age, 
		Dob: pet.Dob,
	}
	return petDto
}