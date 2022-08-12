package usecase

import (
	"log"
	"musement/src/internal/core/contracts"
	"musement/src/internal/core/domain/models"
	"musement/src/internal/core/usecase/internal/formatter"

	"github.com/pkg/errors"
)

type MusementProvider interface {
	GetCities() ([]contracts.City, error)
}

type WeatherProvider interface {
	GetForecastForCity(cityLat, cityLong float64) (contracts.WeatherForecast, error)
}

type forecastedCitiesFormatter interface {
	BuildForecastedCity(city contracts.City, weatherForecast contracts.WeatherForecast) models.CityForecasted
}

type DefaultGetForecastForCitiesUseCase struct {
	musementProvider          MusementProvider
	weatherProvider           WeatherProvider
	forecastedCitiesFormatter forecastedCitiesFormatter
}

func NewDefaultGetForecastForCitiesUseCase(musementProvider MusementProvider,
	weatherProvider WeatherProvider) *DefaultGetForecastForCitiesUseCase {
	return &DefaultGetForecastForCitiesUseCase{
		musementProvider:          musementProvider,
		weatherProvider:           weatherProvider,
		forecastedCitiesFormatter: formatter.NewDefaultForecastedCitiesFormatter(),
	}
}

func (uc *DefaultGetForecastForCitiesUseCase) GetForecastForCities() ([]models.CityForecasted, error) {
	cities, musementErr := uc.musementProvider.GetCities()
	if musementErr != nil {
		return nil, errors.Wrap(musementErr, "error getting cities from musement provider")
	}

	var forecastedCities []models.CityForecasted
	for _, city := range cities {
		log.Printf("getting forecast for city '%s'", city.Name)
		weatherForecast, weatherErr := uc.weatherProvider.GetForecastForCity(city.Latitude, city.Longitude)
		if weatherErr == nil {
			forecastedCity := uc.forecastedCitiesFormatter.BuildForecastedCity(city, weatherForecast)
			forecastedCities = append(forecastedCities, forecastedCity)
		} else {
			log.Print(errors.Wrapf(weatherErr, "error getting forecast for city %s", city.Name))
		}
	}
	return forecastedCities, nil
}
