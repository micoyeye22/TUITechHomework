package getforecast

import "musement/src/internal/core/domain/models"

type GetForecastForCitiesUseCase interface {
	GetForecastForCities() ([]models.CityForecasted, error)
}

type responseFormatter interface {
	FormatCityForcasted(cityForecasted models.CityForecasted) string
}

type responsePrinter interface {
	PrintCity(cityFormatted string)
}
