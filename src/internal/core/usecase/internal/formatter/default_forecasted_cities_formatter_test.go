package formatter

import (
	"musement/src/internal/core/contracts"
	"musement/src/internal/core/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testCityID        = 1
	testCityName      = "test-city"
	testCityLatitude  = 24.60
	testCityLongitude = 50.12
	testConditionText = "test-condition-text"
)

func TestNewDefaultForecastedCitiesFormatter(t *testing.T) {
	expectedFormatter := &defaultForecastedCitiesFormatter{}

	actualFormatter := NewDefaultForecastedCitiesFormatter()

	assert.Equal(t, expectedFormatter, actualFormatter)
}

func TestDefaultForecastedCitiesFormatter_BuildForecastedCity_success(t *testing.T) {
	inputs := []struct {
		description string
		days        int
	}{
		{
			description: "forecast with one day",
			days:        1,
		},
		{
			description: "forecast with two days",
			days:        2,
		},
		{
			description: "forecast with zero days",
			days:        0,
		},
	}

	for _, input := range inputs {
		t.Run(input.description, func(t *testing.T) {
			formatter := &defaultForecastedCitiesFormatter{}

			city := givenACityContract()
			weatherForecast := givenAWeatherForecastContract(input.days)
			expectedCityForecasted := givenAnExpectedCityForecasted(input.days)

			actualCityForecasted := formatter.BuildForecastedCity(city, weatherForecast)

			assert.Equal(t, expectedCityForecasted, actualCityForecasted)
		})
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

func givenAWeatherForecastContract(days int) contracts.WeatherForecast {
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

func givenAnExpectedCityForecasted(days int) models.CityForecasted {
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
