package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/danielperpar/go-pet-api/common"
	"github.com/danielperpar/go-pet-api/domain"
	_ "github.com/lib/pq"
)

type PostgrePetRepositoy struct {
	connectionString string
}

func NewPostgrePetRepositoy(connectionString string) *PostgrePetRepositoy{
	 return &PostgrePetRepositoy{connectionString: connectionString}
}

func (repository *PostgrePetRepositoy)CreatePet(pet domain.Pet) (domain.Pet, error){
	db,errOpen := repository.openConnection()  //TODO: revisar si  abrir/cerrar la conexion por query o si se puede configurar un timeout 
	defer db.Close()

	if errOpen != nil {
		fmt.Println(errOpen)
		return domain.Pet{}, errors.New(common.DbError)
	}

  	sqlStatement := `INSERT INTO pets(name,species,gender,age,dob) VALUES($1,$2,$3,$4,$5)`
   	_,errIns := db.Exec(sqlStatement,pet.Name,pet.Species,pet.Gender,pet.Age,pet.Dob)
	if errIns != nil {
		fmt.Println(errIns)
		return domain.Pet{}, errors.New(common.DbError)
	}

	return pet, nil
}

func (repository *PostgrePetRepositoy)GetPets(start, count int) (*[]domain.Pet, error){
	db,errOpen := repository.openConnection()  //TODO: revisar si abrir/cerrar la conexion por query o si se puede configurar un timeout
	defer db.Close()

	if errOpen != nil {
		fmt.Println(errOpen)
		return nil, errors.New(common.DbError)
	}
	
	sqlStatement := `SELECT id,name,species,gender,age,dob FROM pets LIMIT $1 OFFSET $2`
	rows, errQuery := db.Query(sqlStatement, count, start)
	if errQuery != nil {
		fmt.Println(errQuery)
		return nil, errors.New(common.DbError)
	}
	defer rows.Close()

	pets := []domain.Pet{}

	for rows.Next(){
		var p domain.Pet
		if errScan:= rows.Scan(&p.Id, &p.Name, &p.Species, &p.Gender, &p.Age, &p.Dob); errScan != nil {
			fmt.Println(errScan)
			return nil, errors.New(common.DbError) 
		}
		pets = append(pets, p)
	}

	return &pets, nil
}

func (repository *PostgrePetRepositoy) openConnection() (*sql.DB, error){
	db, err := sql.Open("postgres", repository.connectionString)
	
	if err != nil {
	  fmt.Println(err)
	  return nil, errors.New(common.DbError)
	}
	return db, nil
}

func (repository *PostgrePetRepositoy) ping(db *sql.DB) error {
	fmt.Println("db ping")
	errPing := db.Ping()
	if errPing != nil {
		fmt.Println(errPing)
		return  errors.New(common.DbError)
	}
	fmt.Println("db pong")
	return nil
}