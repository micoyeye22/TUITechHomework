package usecase

import (
	"musement/src/internal/core/contracts"
	"musement/src/internal/core/domain/models"
	"musement/src/internal/core/usecase/internal/formatter"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	testCityID        = 1
	testCityName      = "test-city"
	testCityLatitude  = 24.60
	testCityLongitude = 50.12
	testConditionText = "test-condition-text"
)

func TestNewDefaultGetForecastForCitiesUseCase(t *testing.T) {
	mockMusementProvider, mockWeatherProvider, expectedUseCase :=
		givenADefaultUseCaseWithInjectedComponents()

	actualUseCase := NewDefaultGetForecastForCitiesUseCase(mockMusementProvider, mockWeatherProvider)

	assert.Equal(t, expectedUseCase, actualUseCase)
	thenAssertConstructorMocksExpectations(t, mockMusementProvider, mockWeatherProvider)
}

func TestDefaultGetForecastForCitiesUseCase_GetForecastForCities_whenThereAreCitiesFoundThenReturnsNilError(
	t *testing.T) {
	inputs := []struct {
		description string
		cities      int
		days        int
	}{
		{
			description: "one city scenario with two days forecasted in it",
			cities:      1,
			days:        2,
		},
		{
			description: "two cities scenario with two days forecasted in them",
			cities:      2,
			days:        2,
		},
		{
			description: "two cities scenario with three days forecasted in them",
			cities:      2,
			days:        3,
		},
	}
	for _, input := range inputs {
		t.Run(input.description, func(t *testing.T) {
			mockMusementProvider, mockWeatherProvider, mockForecastedCitiesFormatter, expectedUseCase :=
				givenADefaultUseCaseWithMockedComponents()

			cities := givenACitiesArray(input.cities)
			weatherForecast := givenAWeatherForecast(input.days)

			cityForecasted := givenACityForecasted(input.days)
			expectedForecastedCitiesArray := givenAForecastedCitiesArray(input.cities, cityForecasted)

			mockMusementProvider.On("GetCities").Return(cities, nil)
			mockWeatherProvider.On("GetForecastForCity",
				mock.AnythingOfType("float64"), mock.AnythingOfType("float64")).Return(weatherForecast, nil)
			mockForecastedCitiesFormatter.On("BuildForecastedCity", mock.AnythingOfType("contracts.City"),
				mock.AnythingOfType("contracts.WeatherForecast")).Return(cityForecasted)

			actualForecastedCitiesArray, err := expectedUseCase.GetForecastForCities()

			assert.NoError(t, err)
			assert.Equal(t, expectedForecastedCitiesArray, actualForecastedCitiesArray)
			thenAssertMocksExpectations(t, mockMusementProvider, mockWeatherProvider, mockForecastedCitiesFormatter)
		})
	}
}

func TestDefaultGetForecastForCitiesUseCase_GetForecastForCities_whenThereAreNoCitiesFoundThenReturnsNilError(
	t *testing.T) {
	mockMusementProvider, mockWeatherProvider, mockForecastedCitiesFormatter, expectedUseCase :=
		givenADefaultUseCaseWithMockedComponents()

	cities := givenACitiesArray(0)

	expectedForecastedCitiesArray := givenAForecastedCitiesArray(0, models.CityForecasted{})

	mockMusementProvider.On("GetCities").Return(cities, nil)

	actualForecastedCitiesArray, err := expectedUseCase.GetForecastForCities()

	assert.NoError(t, err)
	assert.Equal(t, expectedForecastedCitiesArray, actualForecastedCitiesArray)
	thenAssertMocksExpectations(t, mockMusementProvider, mockWeatherProvider, mockForecastedCitiesFormatter)
}

func TestDefaultGetForecastForCitiesUseCase_GetForecastForCities_whenMusementProviderFailsThenReturnsError(
	t *testing.T) {
	mockMusementProvider, mockWeatherProvider, mockForecastedCitiesFormatter, expectedUseCase :=
		givenADefaultUseCaseWithMockedComponents()

	musementProviderErr := errors.New("mocked error")
	expectedErrorMessage := "error getting cities from musement provider"

	mockMusementProvider.On("GetCities").Return(nil, musementProviderErr)

	_, err := expectedUseCase.GetForecastForCities()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), musementProviderErr.Error())
	assert.Contains(t, err.Error(), expectedErrorMessage)
	thenAssertMocksExpectations(t, mockMusementProvider, mockWeatherProvider, mockForecastedCitiesFormatter)
}

func TestDefaultGetForecastForCitiesUseCase_GetForecastForCities_whenWeatherProviderFailsForOneThenReturnsEmptyArray(
	t *testing.T) {
	mockMusementProvider, mockWeatherProvider, mockForecastedCitiesFormatter, expectedUseCase :=
		givenADefaultUseCaseWithMockedComponents()

	cities := givenACitiesArray(1)

	weatherErr := errors.New("mocked error")

	mockMusementProvider.On("GetCities").Return(cities, nil)
	mockWeatherProvider.On("GetForecastForCity",
		mock.AnythingOfType("float64"), mock.AnythingOfType("float64")).
		Return(contracts.WeatherForecast{}, weatherErr)

	actualForecastedCitiesArray, err := expectedUseCase.GetForecastForCities()

	assert.NoError(t, err)
	assert.Empty(t, actualForecastedCitiesArray)
	thenAssertMocksExpectations(t, mockMusementProvider, mockWeatherProvider, mockForecastedCitiesFormatter)
}

func givenACityForecasted(days int) models.CityForecasted {
	var forecasts []models.Forecast
	for i := 0; i < days; i++ {
		forecasts = append(forecasts, givenAForecastModel(i))
	}
	return models.CityForecasted{
		Name:      testCityName,
		Forecasts: forecasts,
	}
}

func givenAForecastModel(order int) models.Forecast {
	return models.Forecast{
		Order:       order,
		Description: testConditionText,
	}
}

func givenAForecastedCitiesArray(cities int, cityForecasted models.CityForecasted) []models.CityForecasted {
	var citiesForcastedArray []models.CityForecasted
	for i := 0; i < cities; i++ {
		citiesForcastedArray = append(citiesForcastedArray, cityForecasted)
	}

	return citiesForcastedArray
}

func givenAWeatherForecast(days int) contracts.WeatherForecast {
	var forecastdays []contracts.Forecastday
	for i := 0; i < days; i++ {
		forecastdays = append(forecastdays, givenAForecastday())
	}
	return contracts.WeatherForecast{
		Forecastdays: forecastdays,
	}
}

func givenAForecastday() contracts.Forecastday {
	return contracts.Forecastday{
		Day: givenADay(),
	}
}

func givenADay() contracts.Day {
	return contracts.Day{
		Condition: givenACondition(),
	}
}

func givenACondition() contracts.Condition {
	return contracts.Condition{
		Text: testConditionText,
	}
}

func givenACitiesArray(cities int) []contracts.City {
	var citiesArray []contracts.City
	for i := 0; i < cities; i++ {
		citiesArray = append(citiesArray, givenACity())
	}
	return citiesArray
}

func givenACity() contracts.City {
	return contracts.City{
		ID:        testCityID,
		Name:      testCityName,
		Latitude:  testCityLatitude,
		Longitude: testCityLongitude,
	}
}

func givenADefaultUseCaseWithInjectedComponents() (*MockMusementProvider, *MockWeatherProvider,
	*DefaultGetForecastForCitiesUseCase) {
	mockMusementProvider := new(MockMusementProvider)
	mockWeatherProvider := new(MockWeatherProvider)

	useCase := &DefaultGetForecastForCitiesUseCase{
		musementProvider:          mockMusementProvider,
		weatherProvider:           mockWeatherProvider,
		forecastedCitiesFormatter: formatter.NewDefaultForecastedCitiesFormatter(),
	}

	return mockMusementProvider, mockWeatherProvider, useCase
}

func thenAssertConstructorMocksExpectations(t *testing.T, mockMusementProvider *MockMusementProvider,
	mockWeatherProvider *MockWeatherProvider) {
	mockMusementProvider.AssertExpectations(t)
	mockWeatherProvider.AssertExpectations(t)
}

func givenADefaultUseCaseWithMockedComponents() (*MockMusementProvider, *MockWeatherProvider,
	*MockForecastedCitiesFormatter, *DefaultGetForecastForCitiesUseCase) {
	mockMusementProvider := new(MockMusementProvider)
	mockWeatherProvider := new(MockWeatherProvider)
	mockForecastedCitiesFormatter := new(MockForecastedCitiesFormatter)

	useCase := &DefaultGetForecastForCitiesUseCase{
		musementProvider:          mockMusementProvider,
		weatherProvider:           mockWeatherProvider,
		forecastedCitiesFormatter: mockForecastedCitiesFormatter,
	}

	return mockMusementProvider, mockWeatherProvider, mockForecastedCitiesFormatter, useCase
}

func thenAssertMocksExpectations(t *testing.T, mockMusementProvider *MockMusementProvider,
	mockWeatherProvider *MockWeatherProvider, mockForecastedCitiesFormatter *MockForecastedCitiesFormatter) {
	mockMusementProvider.AssertExpectations(t)
	mockWeatherProvider.AssertExpectations(t)
	mockForecastedCitiesFormatter.AssertExpectations(t)
}
