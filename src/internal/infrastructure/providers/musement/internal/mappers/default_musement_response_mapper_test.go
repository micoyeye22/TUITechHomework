package mappers

import (
	"musement/src/internal/core/contracts"
	"musement/src/internal/infrastructure/providers/musement/internal/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testCityID        = 1
	testCityName      = "test-city"
	testCityLatitude  = 24.60
	testCityLongitude = 50.12
)

func TestNewDefaultMusementResponseMapper(t *testing.T) {
	expectedMapper := &defaultMusementResponseMapper{}

	actualMapper := NewDefaultMusementResponseMapper()

	assert.Equal(t, expectedMapper, actualMapper)
}

func TestDefaultMusementResponseMapper_ToCityContract_success(t *testing.T) {
	expectedMapper := &defaultMusementResponseMapper{}

	responseCity := givenAResponseCity()
	expectedCity := givenAContractCity()

	actualCity := expectedMapper.ToCityContract(responseCity)

	assert.Equal(t, expectedCity, actualCity)
}

func givenAResponseCity() response.City {
	return response.City{
		ID:        testCityID,
		Name:      testCityName,
		Latitude:  testCityLatitude,
		Longitude: testCityLongitude,
	}
}

func givenAContractCity() contracts.City {
	return contracts.City{
		ID:        testCityID,
		Name:      testCityName,
		Latitude:  testCityLatitude,
		Longitude: testCityLongitude,
	}
}
