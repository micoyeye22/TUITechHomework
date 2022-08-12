package mappers

import (
	"musement/src/internal/core/contracts"
	"musement/src/internal/infrastructure/providers/weather/internal/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testConditionText = "Sunny"
	days              = 2
)

func TestNewDefaultWeatherResponseMapper(t *testing.T) {
	expectedMapper := &defaultWeatherResponseMapper{}

	actualMapper := NewDefaultWeatherResponseMapper()

	assert.Equal(t, expectedMapper, actualMapper)
}

func TestDefaultWeatherResponseMapper_ToWeatherForecastContract_success(t *testing.T) {
	mapper := &defaultWeatherResponseMapper{}

	weatherAPIResponse := givenAWeatherAPIResponse()
	expectedWeatherForecast := givenAWeatherForecastContract()

	actualWeatherForecast := mapper.ToWeatherForecastContract(weatherAPIResponse)

	assert.Equal(t, expectedWeatherForecast, actualWeatherForecast)
}

func givenAWeatherForecastContract() contracts.WeatherForecast {
	return contracts.WeatherForecast{
		Forecastdays: givenAContractForecastdaysArray(),
	}
}

func givenAContractForecastdaysArray() []contracts.Forecastday {
	var forecastdays []contracts.Forecastday
	for i := 0; i < days; i++ {
		forecastdays = append(forecastdays, givenAContractForecastday())
	}
	return forecastdays
}

func givenAContractForecastday() contracts.Forecastday {
	return contracts.Forecastday{
		Day: contracts.Day{
			Condition: contracts.Condition{
				Text: testConditionText,
			},
		},
	}
}

func givenAWeatherAPIResponse() response.WeatherAPIResponse {
	return response.WeatherAPIResponse{
		Forecast: givenAResponseForecast(),
	}
}

func givenAResponseForecast() response.Forecast {
	return response.Forecast{
		Forecastdays: givenAResponseForecastdaysArray(),
	}
}

func givenAResponseForecastdaysArray() []response.Forecastday {
	var forecastdaysArray []response.Forecastday
	for i := 0; i < days; i++ {
		forecastdaysArray = append(forecastdaysArray, givenAResponseForecastday())
	}
	return forecastdaysArray
}

func givenAResponseForecastday() response.Forecastday {
	return response.Forecastday{
		Day: response.Day{
			Condition: response.Condition{
				Text: testConditionText,
			},
		},
	}
}
