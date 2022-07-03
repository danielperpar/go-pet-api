package infrastructure

import (
	"database/sql"
	"errors"
	"github.com/danielperpar/go-pet-api/common"
	"github.com/danielperpar/go-pet-api/domain"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

type PostgrePetRepositoy struct {
}

func NewPostgrePetRepositoy() *PostgrePetRepositoy{
	 return &PostgrePetRepositoy{}
}

func (repository *PostgrePetRepositoy) OpenConnection() error{
	connString := GetConnectionString()
	var err error
	db, err = sql.Open("postgres", connString)
	return err
}

func (repository *PostgrePetRepositoy)CreatePet(pet domain.Pet) (domain.Pet, error){
  	sqlStatement := `INSERT INTO pets(name,species,gender,age,dob) VALUES($1,$2,$3,$4,$5)`
   	_,errIns := db.Exec(sqlStatement,pet.Name,pet.Species,pet.Gender,pet.Age,pet.Dob)
	if errIns != nil {
		log.Println(errIns)
		return domain.Pet{}, errors.New(common.DbError)
	}

	return pet, nil
}

func (repository *PostgrePetRepositoy)GetPets(start, count int) (*[]domain.Pet, error){
	sqlStatement := `SELECT id,name,species,gender,age,dob FROM pets LIMIT $1 OFFSET $2`
	rows, errQuery := db.Query(sqlStatement, count, start)
	if errQuery != nil {
		log.Println(errQuery)
		return nil, errors.New(common.DbError)
	}
	defer rows.Close()

	pets := []domain.Pet{}

	for rows.Next(){
		var p domain.Pet
		if errScan:= rows.Scan(&p.Id, &p.Name, &p.Species, &p.Gender, &p.Age, &p.Dob); errScan != nil {
			log.Println(errScan)
			return nil, errors.New(common.DbError) 
		}
		pets = append(pets, p)
	}

	return &pets, nil
}