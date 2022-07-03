package infrastructure

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/danielperpar/go-pet-api/common"
	"github.com/danielperpar/go-pet-api/domain"

	_ "github.com/lib/pq"
)

var db *sql.DB

type PostgrePetRepositoy struct {
}

func NewPostgrePetRepositoy() *PostgrePetRepositoy{
	 return &PostgrePetRepositoy{}
}

func (repository *PostgrePetRepositoy) OpenConnection() error{
	connString := GetConnectionString()
	var errOpen error
	db, errOpen = sql.Open("postgres", connString)
	if(errOpen != nil){
		return common.NewError(http.StatusInternalServerError, errOpen.Error())
	}
	return nil
}

func (repository *PostgrePetRepositoy)CreatePet(pet domain.Pet) (domain.Pet, error){
  	sqlStatement := `INSERT INTO pets(name,species,gender,age,dob) VALUES($1,$2,$3,$4,$5)`
   	_,errIns := db.Exec(sqlStatement,pet.Name,pet.Species,pet.Gender,pet.Age,pet.Dob)
	if errIns != nil {
		log.Println(errIns)
		return domain.Pet{}, common.NewError(http.StatusInternalServerError, errIns.Error())
	}
	return pet, nil
}

func (repository *PostgrePetRepositoy)GetPets() (*[]domain.Pet, error){
	sqlStatement := `SELECT id,name,species,gender,age,dob FROM pets`
	rows, errQuery := db.Query(sqlStatement)
	if errQuery != nil {
		log.Println(errQuery)
		return nil, common.NewError(http.StatusInternalServerError, errQuery.Error())
	}
	defer rows.Close()
	pets := []domain.Pet{}
	for rows.Next(){
		var p domain.Pet
		if errScan:= rows.Scan(&p.Id, &p.Name, &p.Species, &p.Gender, &p.Age, &p.Dob); errScan != nil {
			log.Println(errScan)
			return nil, common.NewError(http.StatusInternalServerError, errScan.Error())
		}
		pets = append(pets, p)
	}
	return &pets, nil
}