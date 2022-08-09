package formatter

import (
	"musement/src/internal/core/contracts"
	"musement/src/internal/core/domain/models"
)

type defaultForecastedCitiesFormatter struct {
}

func NewDefaultForecastedCitiesFormatter() *defaultForecastedCitiesFormatter {
	return &defaultForecastedCitiesFormatter{}
}

func (f *defaultForecastedCitiesFormatter) BuildForecastedCity(city contracts.City,
	weatherForecast contracts.WeatherForecast) models.CityForecasted {
	cityForecasted := models.CityForecasted{
		Name: city.Name,
	}
	for index, forecastday := range weatherForecast.Forecastdays {
		forecast := models.Forecast{
			Order:       index,
			Description: forecastday.Day.Condition.Text,
		}
		cityForecasted.Forecasts = append(cityForecasted.Forecasts, forecast)
	}
	return cityForecasted
}
