package client

import (
	"encoding/json"
	"fmt"
	"io"
	"musement/src/internal/infrastructure/providers/weather/config"
	"musement/src/internal/infrastructure/providers/weather/internal/response"
	"net/http"

	"github.com/pkg/errors"
)

type Config interface {
	WeatherProviderClientConfig() config.WeatherProviderClientConfig
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type defaultRestWeatherClient struct {
	config     Config
	httpClient HTTPClient
}

func NewDefaultRestWeatherClient(config Config, httpClient HTTPClient) *defaultRestWeatherClient {
	return &defaultRestWeatherClient{
		config:     config,
		httpClient: httpClient,
	}
}

func (rc *defaultRestWeatherClient) GetForecastForCityRequest(cityLat, cityLong float64) (
	response.WeatherAPIResponse, error) {
	baseURL := rc.config.WeatherProviderClientConfig().BaseURL

	url := fmt.Sprintf("%s/forecast.json?key=%s&q=%v,%v&days=%d",
		baseURL, rc.config.WeatherProviderClientConfig().WeatherClientToken,
		cityLat, cityLong, rc.config.WeatherProviderClientConfig().ForecastDays)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		apiError := errors.Wrapf(err,
			"error building request to get forecast for city with latitude %v and longitude %v", cityLat, cityLong)
		return response.WeatherAPIResponse{}, apiError
	}

	responses, err := rc.httpClient.Do(req)
	if err != nil {
		apiError := errors.Wrapf(err,
			"error doing request to get forecast for city with latitude %v and longitude %v", cityLat, cityLong)
		return response.WeatherAPIResponse{}, apiError
	}
	defer func() { _ = responses.Body.Close() }()

	if responses.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(responses.Body)
		apiError := errors.Errorf("invalid status in response getting forecast code '%d' and body '%s'",
			responses.StatusCode, body)
		return response.WeatherAPIResponse{}, apiError
	}

	var weatherAPIResponse response.WeatherAPIResponse
	decoder := json.NewDecoder(responses.Body)
	err = decoder.Decode(&weatherAPIResponse)

	return weatherAPIResponse, err
}
