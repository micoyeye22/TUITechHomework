package client

import (
	"encoding/json"
	"fmt"
	"musement/src/internal/infrastructure/providers/weather/config"
	"musement/src/internal/infrastructure/providers/weather/internal/response"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testBaseURL        = "host://api"
	testInvalidBaseURL = "$%&!invalidURL/!%$&"
	testCityLatitude   = 48.5
	testCityLongitude  = 20.1
	testToken          = "asd123"
	testForecastDays   = 2
	testConditionText  = "Sunny"
)

func TestNewDefaultRestWeatherClient(t *testing.T) {
	httpClient, mockConfig, expectedRestWeatherClient := givenADefaultRestWeatherClientWithMockedComponents()

	actualRestWeatherClient := NewDefaultRestWeatherClient(mockConfig, httpClient)

	assert.Equal(t, expectedRestWeatherClient, actualRestWeatherClient)
	mockConfig.AssertExpectations(t)
}

func TestDefaultRestWeatherClient_GetForecastForCityRequest_success(t *testing.T) {
	_, mockConfig, restWeatherClient := givenADefaultRestWeatherClientWithMockedComponents()
	weatherRestClientConfig := givenAWeatherRestClientConfig(testBaseURL)

	expectedWeatherAPIResponse := givenAGetForecastForCityRequestResponse()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockConfig.On("WeatherProviderClientConfig").Return(weatherRestClientConfig)
	whenGetForecastForCityRequestReturnsResponse(http.StatusOK, expectedWeatherAPIResponse)

	actualWeatherAPIResponse, err := restWeatherClient.GetForecastForCityRequest(testCityLatitude, testCityLongitude)

	assert.NoError(t, err)
	assert.Equal(t, expectedWeatherAPIResponse, actualWeatherAPIResponse)
	mockConfig.AssertExpectations(t)
}

func TestDefaultRestWeatherClient_GetForecastForCityRequest_whenBuildingRequestFailsThenReturnsError(t *testing.T) {
	_, mockConfig, restWeatherClient := givenADefaultRestWeatherClientWithMockedComponents()
	weatherRestClientConfig := givenAWeatherRestClientConfig(testInvalidBaseURL)

	expectedErrorMessage := fmt.Sprintf(
		"error building request to get forecast for city with latitude %v and longitude %v",
		testCityLatitude, testCityLongitude)

	mockConfig.On("WeatherProviderClientConfig").Return(weatherRestClientConfig)

	_, err := restWeatherClient.GetForecastForCityRequest(testCityLatitude, testCityLongitude)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
	mockConfig.AssertExpectations(t)
}

func TestDefaultRestWeatherClient_GetForecastForCityRequest_whenClientFailsThenReturnsError(t *testing.T) {
	_, mockConfig, restWeatherClient := givenADefaultRestWeatherClientWithMockedComponents()
	weatherRestClientConfig := givenAWeatherRestClientConfig(testBaseURL)

	expectedErrorMessage := fmt.Sprintf(
		"error doing request to get forecast for city with latitude %v and longitude %v",
		testCityLatitude, testCityLongitude)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mockConfig.On("WeatherProviderClientConfig").Return(weatherRestClientConfig)

	_, err := restWeatherClient.GetForecastForCityRequest(testCityLatitude, testCityLongitude)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
	mockConfig.AssertExpectations(t)
}

func TestDefaultRestWeatherClient_GetForecastForCityRequest_whenResponseHasInvalidStatusThenReturnsError(t *testing.T) {
	_, mockConfig, restWeatherClient := givenADefaultRestWeatherClientWithMockedComponents()
	weatherRestClientConfig := givenAWeatherRestClientConfig(testBaseURL)

	httpStatus := http.StatusBadRequest
	responseStruct := response.WeatherAPIResponse{}
	responseBody := givenAResponseBodyString(t, responseStruct)

	expectedErrorMessage := fmt.Sprintf("invalid status in response getting forecast code '%d' and body '%s'",
		httpStatus, responseBody)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockConfig.On("WeatherProviderClientConfig").Return(weatherRestClientConfig)
	whenGetForecastForCityRequestReturnsResponse(httpStatus, responseStruct)

	_, err := restWeatherClient.GetForecastForCityRequest(testCityLatitude, testCityLongitude)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
	mockConfig.AssertExpectations(t)
}

func givenAResponseBodyString(t *testing.T, weatherAPIResponse response.WeatherAPIResponse) string {
	responseBodyBytes, marshallErr := json.Marshal(weatherAPIResponse)
	require.NoError(t, marshallErr)
	return string(responseBodyBytes)
}

func givenAWeatherRestClientConfig(url string) config.WeatherProviderClientConfig {
	return config.WeatherProviderClientConfig{
		BaseURL:            url,
		WeatherClientToken: testToken,
		ForecastDays:       testForecastDays,
	}
}

func whenGetForecastForCityRequestReturnsResponse(status int, weatherAPIResponse response.WeatherAPIResponse) {
	httpMockURL := fmt.Sprintf("%s/forecast.json?key=%s&q=%v,%v&days=%d",
		testBaseURL, testToken, testCityLatitude, testCityLongitude, testForecastDays)

	httpmock.RegisterResponder(
		http.MethodGet,
		httpMockURL,
		func(req *http.Request) (resp *http.Response, e error) {
			resp, err := httpmock.NewJsonResponse(status, weatherAPIResponse)
			return resp, err
		})
}

func givenAGetForecastForCityRequestResponse() response.WeatherAPIResponse {
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
	for i := 0; i < testForecastDays; i++ {
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

func givenADefaultRestWeatherClientWithMockedComponents() (*http.Client, *MockConfig, *defaultRestWeatherClient) {
	httpClient := http.DefaultClient
	mockConfig := new(MockConfig)

	restWeatherClient := &defaultRestWeatherClient{
		config:     mockConfig,
		httpClient: httpClient,
	}

	return httpClient, mockConfig, restWeatherClient
}
