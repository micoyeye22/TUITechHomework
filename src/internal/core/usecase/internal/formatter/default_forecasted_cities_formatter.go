package formatter

import (
	"musement/src/internal/core/contracts"
	"musement/src/internal/core/domain/entities"
)

type defaultForecastedCitiesFormatter struct {
}

func NewDefaultForecastedCitiesFormatter() *defaultForecastedCitiesFormatter {
	return &defaultForecastedCitiesFormatter{}
}

func (f *defaultForecastedCitiesFormatter) BuildForecastedCity(city contracts.City,
	weatherForecast contracts.WeatherForecast) entities.CityForecasted {
	cityForecasted := entities.CityForecasted{
		Name: city.Name,
	}
	for index, forecastday := range weatherForecast.Forecastdays {
		forecast := entities.Forecast{
			Order:       index,
			Description: forecastday.Day.Condition.Text,
		}
		cityForecasted.Forecasts = append(cityForecasted.Forecasts, forecast)
	}
	return cityForecasted
}
