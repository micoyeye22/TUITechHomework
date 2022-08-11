package weather

import (
	"musement/src/internal/core/contracts"
	"musement/src/internal/infrastructure/providers/weather/internal/client"
	"musement/src/internal/infrastructure/providers/weather/internal/mappers"
	"musement/src/internal/infrastructure/providers/weather/internal/response"
	"net/http"

	"github.com/pkg/errors"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WeatherProviderClientConfig interface {
	BaseURL() string
	WeatherClientToken() string
	ForecastDays() int
}

type restWeatherClient interface {
	GetForecastForCityRequest(cityLat, cityLong float64) (response.WeatherAPIResponse, error)
}

type weatherResponseMapper interface {
	ToWeatherForecastContract(weatherAPIResponse response.WeatherAPIResponse) contracts.WeatherForecast
}

type DefaultWeatherProvider struct {
	restWeatherClient restWeatherClient
	mapper            weatherResponseMapper
}

func NewDefaultWeatherProvider(config WeatherProviderClientConfig, httpClient HTTPClient) *DefaultWeatherProvider {
	return &DefaultWeatherProvider{
		restWeatherClient: client.NewDefaultRestWeatherClient(config, httpClient),
		mapper:            mappers.NewDefaultWeatherResponseMapper(),
	}
}

func (p *DefaultWeatherProvider) GetForecastForCity(cityLat, cityLong float64) (contracts.WeatherForecast, error) {
	weatherAPIResponse, reqErr := p.restWeatherClient.GetForecastForCityRequest(cityLat, cityLong)
	if reqErr != nil {
		return contracts.WeatherForecast{}, errors.Wrap(reqErr, "error making request to weatherAPI")
	}

	weatherForecastContract := p.mapper.ToWeatherForecastContract(weatherAPIResponse)

	return weatherForecastContract, nil
}
