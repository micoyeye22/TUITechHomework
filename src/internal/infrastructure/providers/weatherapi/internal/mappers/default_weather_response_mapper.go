package mappers

import (
	"musement/src/internal/core/contracts"
	"musement/src/internal/infrastructure/providers/weatherapi/internal/response"
)

type defaultWeatherResponseMapper struct {
}

func NewDefaultWeatherResponseMapper() *defaultWeatherResponseMapper {
	return &defaultWeatherResponseMapper{}
}

func (m *defaultWeatherResponseMapper) ToWeatherForecastContract(
	weatherAPIResponse response.WeatherAPIResponse) contracts.WeatherForecast {
	var forecastdays []contracts.Forecastday
	for _, forecastday := range weatherAPIResponse.Forecast.Forecastdays {
		forecastdays = append(forecastdays, m.toForecastdayContract(forecastday))
	}
	return contracts.WeatherForecast{
		Forecastdays: forecastdays,
	}
}

func (m *defaultWeatherResponseMapper) toForecastdayContract(
	responseForecastday response.Forecastday) contracts.Forecastday {
	return contracts.Forecastday{
		Day: m.toDayContract(responseForecastday.Day),
	}
}

func (m *defaultWeatherResponseMapper) toDayContract(responseDay response.Day) contracts.Day {
	return contracts.Day{
		Condition: m.toConditionContract(responseDay.Condition),
	}
}

func (m *defaultWeatherResponseMapper) toConditionContract(responseCondition response.Condition) contracts.Condition {
	return contracts.Condition{
		Text: responseCondition.Text,
	}
}
