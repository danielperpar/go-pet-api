package domain

import (
	"testing"
	"time"
	"net/http"

	"github.com/danielperpar/go-pet-api/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) CreatePet(pet *Pet) error {
	args := mock.Called()
	result := args.Get(0)
	return result.(error)
}

func (mock *MockRepository) GetPets() (*[]Pet,error){
	args := mock.Called()
	result := args.Get(0)
	return result.(*[]Pet), args.Error(1)
}

func TestGetKpi_ok(t *testing.T){

	//Arrange
	testData := []Pet{
		*NewPet("tom", "cat", "male", 1, time.Now()),
		*NewPet("kiut", "cat", "male", 1, time.Now()),
		*NewPet("trek", "cat", "male", 1, time.Now()),
		*NewPet("tor", "dog", "male", 1, time.Now()),
	}

	mockRepo := new (MockRepository)
	mockRepo.On("GetPets").Return(&testData, nil)
	service := NewStatsService(mockRepo)

	//Act
	kpi, err := service.GetKpi("cat")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, float32(1), kpi.AvgAge)
	sp := *kpi.PredomSpec
	assert.Equal(t, "cat", sp[0])
	assert.Equal(t, nil, err)
}

func TestGetKpi_species_not_found(t *testing.T){

	//Arrange
	testData := []Pet{
		*NewPet("tom", "cat", "male", 1, time.Now()),
		*NewPet("kiut", "cat", "male", 1, time.Now()),
		*NewPet("trek", "cat", "male", 1, time.Now()),
		*NewPet("tor", "dog", "male", 1, time.Now()),
	}

	mockRepo := new (MockRepository)
	mockRepo.On("GetPets").Return(&testData, nil)
	service := NewStatsService(mockRepo)

	//Act
	_, err := service.GetKpi("not-existent")

	//Assert
	mockRepo.AssertExpectations(t)
	customErr := err.(*common.Error)
	 assert.Equal(t, http.StatusNotFound, customErr.Code)
	 assert.Equal(t, common.Domain_PetNotFound, customErr.Message)
}

