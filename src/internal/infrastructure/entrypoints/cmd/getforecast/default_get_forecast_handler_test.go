package getforecast

import (
	"github.com/pkg/errors"
	"musement/src/internal/core/domain/models"
	"musement/src/internal/infrastructure/entrypoints/cmd/getforecast/internal/formatters"
	"musement/src/internal/infrastructure/entrypoints/cmd/getforecast/internal/printers"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testCityName      = "Berlin"
	testForecast      = "Sunny"
	testCityFormatted = "Processed city Berlin | Sunny"
)

func TestNewDefaultGetForecastHandler(t *testing.T) {
	mockUseCase, expectedGetForecastHandler := givenAGetForecastHandlerWithInjectedComponents()

	actualGetForecastHandler := NewDefaultGetForecastHandler(mockUseCase)

	assert.Equal(t, expectedGetForecastHandler, actualGetForecastHandler)
	mockUseCase.AssertExpectations(t)
}

func TestDefaultGetForecastHandler_HandleGetForecast_whenUseCaseReturnsCitiesSuccessfullyThenReturnsNilError(
	t *testing.T) {
	mockUseCase, mockResponseFormatter, mockResponsePrinter, getForecastHandler :=
		givenAGetForecastHandlerWithMockedComponents()

	citiesForecasted := givenACityForecastArray()

	mockUseCase.On("GetForecastForCities").Return(citiesForecasted, nil)
	mockResponseFormatter.On("FormatCityForcasted", citiesForecasted[0]).Return(testCityFormatted)
	mockResponsePrinter.On("PrintCity", testCityFormatted).Return()

	err := getForecastHandler.HandleGetForecast()

	assert.NoError(t, err)
	thenAssertMockExpectations(t, mockUseCase, mockResponseFormatter, mockResponsePrinter)
}

func TestDefaultGetForecastHandler_HandleGetForecast_whenUseCaseFailsThenReturnsError(t *testing.T) {
	mockUseCase, mockResponseFormatter, mockResponsePrinter, getForecastHandler :=
		givenAGetForecastHandlerWithMockedComponents()

	useCaseErr := errors.New("error in useCase")
	expectedErrorMessage := "error in getForecast processing"
	var citiesForecasted []models.CityForecasted

	mockUseCase.On("GetForecastForCities").Return(citiesForecasted, useCaseErr)

	err := getForecastHandler.HandleGetForecast()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), useCaseErr.Error())
	assert.Contains(t, err.Error(), expectedErrorMessage)
	thenAssertMockExpectations(t, mockUseCase, mockResponseFormatter, mockResponsePrinter)
}

func givenACityForecastArray() []models.CityForecasted {
	return []models.CityForecasted{
		{
			Name: testCityName,
			Forecasts: []models.Forecast{
				{
					Order:       0,
					Description: testForecast,
				},
			},
		},
	}
}

func thenAssertMockExpectations(t *testing.T, mockGetForecastForCitiesUseCase *MockGetForecastForCitiesUseCase,
	mockResponseFormatter *MockResponseFormatter, mockResponsePrinter *MockResponsePrinter) {
	mockGetForecastForCitiesUseCase.AssertExpectations(t)
	mockResponseFormatter.AssertExpectations(t)
	mockResponsePrinter.AssertExpectations(t)
}

func givenAGetForecastHandlerWithMockedComponents() (*MockGetForecastForCitiesUseCase, *MockResponseFormatter,
	*MockResponsePrinter, *DefaultGetForecastHandler) {
	mockUseCase := new(MockGetForecastForCitiesUseCase)
	mockResponseFormatter := new(MockResponseFormatter)
	mockResponsePrinter := new(MockResponsePrinter)

	handler := &DefaultGetForecastHandler{
		GetForecastForCitiesUseCase: mockUseCase,
		formatter:                   mockResponseFormatter,
		printer:                     mockResponsePrinter,
	}

	return mockUseCase, mockResponseFormatter, mockResponsePrinter, handler
}

func givenAGetForecastHandlerWithInjectedComponents() (*MockGetForecastForCitiesUseCase, *DefaultGetForecastHandler) {
	mockUseCase := new(MockGetForecastForCitiesUseCase)

	handler := &DefaultGetForecastHandler{
		GetForecastForCitiesUseCase: mockUseCase,
		formatter:                   formatters.NewDefaultResponseFormatter(),
		printer:                     printers.NewDefaultResponsePrinter(),
	}

	return mockUseCase, handler
}
