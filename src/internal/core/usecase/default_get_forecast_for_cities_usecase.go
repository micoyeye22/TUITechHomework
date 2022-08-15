package usecase

import (
	"log"
	"musement/src/internal/core/contracts"
	"musement/src/internal/core/domain/entities"
	"musement/src/internal/core/usecase/internal/formatter"
	"sync"

	"github.com/pkg/errors"
)

type MusementProvider interface {
	GetCities() ([]contracts.City, error)
}

type WeatherProvider interface {
	GetForecastForCity(cityLat, cityLong float64) (contracts.WeatherForecast, error)
}

type forecastedCitiesFormatter interface {
	BuildForecastedCity(city contracts.City, weatherForecast contracts.WeatherForecast) entities.CityForecasted
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

func (uc *DefaultGetForecastForCitiesUseCase) GetForecastForCities() ([]entities.CityForecasted, error) {
	cities, musementErr := uc.musementProvider.GetCities()
	if musementErr != nil {
		return nil, errors.Wrap(musementErr, "error getting cities from musement provider")
	}

	var wg sync.WaitGroup
	var forecastedCities []entities.CityForecasted

	for _, city := range cities {
		wg.Add(1)
		go func(city contracts.City) {
			defer wg.Done()
			weatherForecast, weatherErr := uc.weatherProvider.GetForecastForCity(city.Latitude, city.Longitude)
			if weatherErr == nil {
				forecastedCity := uc.forecastedCitiesFormatter.BuildForecastedCity(city, weatherForecast)
				forecastedCities = append(forecastedCities, forecastedCity)
			} else {
				log.Print(errors.Wrapf(weatherErr, "error getting forecast for city %s", city.Name))
			}
		}(city)
	}
	wg.Wait()
	return forecastedCities, nil
}
