package musement

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"musement/src/internal/core/contracts"
	"musement/src/internal/infrastructure/providers/musement/internal/client"
	"musement/src/internal/infrastructure/providers/musement/internal/mappers"
	"musement/src/internal/infrastructure/providers/musement/internal/response"
	"testing"
)

const (
	testCityID        = 1
	testCityName      = "test-city"
	testCityLatitude  = 24.60
	testCityLongitude = 50.12
)

func TestNewDefaultMusementProvider(t *testing.T) {
	mockConfig, mockHTTPClient, expectedMusementProvider := givenADefaultMusementProviderWithInjectedComponents()

	actualMusementProvider := NewDefaultMusementProvider(mockConfig, mockHTTPClient)

	assert.Equal(t, expectedMusementProvider, actualMusementProvider)
	mockConfig.AssertExpectations(t)
	mockHTTPClient.AssertExpectations(t)
}

func TestDefaultMusementProvider_GetCities_whenCitiesAreFoundThenReturnsArray(t *testing.T) {
	inputs := []struct {
		description string
		cities      int
	}{
		{
			description: "one cities scenario",
			cities:      1,
		},
		{
			description: "two cities scenario",
			cities:      2,
		},
	}
	for _, input := range inputs {
		t.Run(input.description, func(t *testing.T) {
			mockConfig, mockHTTPClient, mockRestMusementClient, mockMusementResponseMapper, musementProvider :=
				givenADefaultMusementProviderWithMockedComponents()

			musementAPIResponse := givenAMusementAPIResponse(input.cities)
			cityContract := givenACityContract()
			expectedCitiesContractArray := givenACitiesContractArray(input.cities, cityContract)

			mockRestMusementClient.On("GetCitiesRequest").Return(musementAPIResponse, nil)
			mockMusementResponseMapper.On("ToCityContract", mock.AnythingOfType("response.City")).
				Return(cityContract)

			actualCitiesContractArray, err := musementProvider.GetCities()

			assert.NoError(t, err)
			assert.Equal(t, expectedCitiesContractArray, actualCitiesContractArray)
			thenAssertMocksExpectations(t, mockConfig, mockHTTPClient, mockRestMusementClient,
				mockMusementResponseMapper)
		})
	}
}

func TestDefaultMusementProvider_GetCities_whenZeroCitiesAreFoundThenReturnsEmptyArray(t *testing.T) {
	mockConfig, mockHTTPClient, mockRestMusementClient, mockMusementResponseMapper, musementProvider :=
		givenADefaultMusementProviderWithMockedComponents()

	musementAPIResponse := givenAMusementAPIResponse(0)

	mockRestMusementClient.On("GetCitiesRequest").Return(musementAPIResponse, nil)

	actualCitiesContractArray, err := musementProvider.GetCities()

	assert.NoError(t, err)
	assert.Empty(t, actualCitiesContractArray)
	thenAssertMocksExpectations(t, mockConfig, mockHTTPClient, mockRestMusementClient,
		mockMusementResponseMapper)
}

func TestDefaultMusementProvider_GetCities_whenRestClientFailsThenReturnsError(t *testing.T) {
	mockConfig, mockHTTPClient, mockRestMusementClient, mockMusementResponseMapper, musementProvider :=
		givenADefaultMusementProviderWithMockedComponents()

	musementAPIResponse := givenAMusementAPIResponse(0)
	restClientErr := errors.New("error in rest client")
	expectedErrorMessage := "error making request to musementAPI"

	mockRestMusementClient.On("GetCitiesRequest").Return(musementAPIResponse, restClientErr)

	_, err := musementProvider.GetCities()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), restClientErr.Error())
	assert.Contains(t, err.Error(), expectedErrorMessage)
	thenAssertMocksExpectations(t, mockConfig, mockHTTPClient, mockRestMusementClient,
		mockMusementResponseMapper)
}

func givenACitiesContractArray(cities int, contract contracts.City) []contracts.City {
	var cityContractArray []contracts.City
	for i := 0; i < cities; i++ {
		cityContractArray = append(cityContractArray, contract)
	}
	return cityContractArray
}

func givenAMusementAPIResponse(cities int) response.MusementAPIResponse {
	return response.MusementAPIResponse{
		Cities: givenACitiesArray(cities),
	}
}

func givenACitiesArray(cities int) []response.City {
	var citiesArray []response.City
	for i := 0; i < cities; i++ {
		citiesArray = append(citiesArray, givenACity())
	}
	return citiesArray
}

func givenACity() response.City {
	return response.City{
		ID:        testCityID,
		Name:      testCityName,
		Latitude:  testCityLatitude,
		Longitude: testCityLongitude,
	}
}

func givenACityContract() contracts.City {
	return contracts.City{
		ID:        testCityID,
		Name:      testCityName,
		Latitude:  testCityLatitude,
		Longitude: testCityLongitude,
	}
}

func givenADefaultMusementProviderWithMockedComponents() (*MockConfig, *MockHTTPClient,
	*MockRestMusementClient, *MockMusementResponseMapper, *DefaultMusementProvider) {
	mockConfig := new(MockConfig)
	mockHTTPClient := new(MockHTTPClient)
	mockRestMusementClient := new(MockRestMusementClient)
	mockMusementResponseMapper := new(MockMusementResponseMapper)

	musementProvider := &DefaultMusementProvider{
		restMusementClient: mockRestMusementClient,
		mapper:             mockMusementResponseMapper,
	}

	return mockConfig, mockHTTPClient, mockRestMusementClient, mockMusementResponseMapper, musementProvider
}

func givenADefaultMusementProviderWithInjectedComponents() (*MockConfig, *MockHTTPClient,
	*DefaultMusementProvider) {
	mockConfig := new(MockConfig)
	mockHTTPClient := new(MockHTTPClient)

	musementProvider := &DefaultMusementProvider{
		restMusementClient: client.NewDefaultRestMusementClient(mockConfig, mockHTTPClient),
		mapper:             mappers.NewDefaultMusementResponseMapper(),
	}

	return mockConfig, mockHTTPClient, musementProvider
}

func thenAssertMocksExpectations(t *testing.T, mockConfig *MockConfig,
	mockHTTPClient *MockHTTPClient, mockRestMusementClient *MockRestMusementClient,
	mockMusementResponseMapper *MockMusementResponseMapper) {
	mockConfig.AssertExpectations(t)
	mockHTTPClient.AssertExpectations(t)
	mockRestMusementClient.AssertExpectations(t)
	mockMusementResponseMapper.AssertExpectations(t)
}
