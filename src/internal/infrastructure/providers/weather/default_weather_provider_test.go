package weather

import (
	"errors"
	"musement/src/internal/core/contracts"
	"musement/src/internal/infrastructure/providers/weather/internal/client"
	"musement/src/internal/infrastructure/providers/weather/internal/mappers"
	"musement/src/internal/infrastructure/providers/weather/internal/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testConditionText = "Sunny"
	testCityLatitude  = 24.60
	testCityLongitude = 50.12
)

func TestNewDefaultWeatherProvider(t *testing.T) {
	mockConfig, mockHTTPClient, expectedWeatherProvider := givenADefaultWeatherProviderWithInjectedComponents()

	actualWeatherProvider := NewDefaultWeatherProvider(mockConfig, mockHTTPClient)

	assert.Equal(t, expectedWeatherProvider, actualWeatherProvider)
	mockConfig.AssertExpectations(t)
	mockHTTPClient.AssertExpectations(t)
}

func TestDefaultWeatherProvider_GetForecastForCity_success(t *testing.T) {
	mockConfig, mockHTTPClient, mockRestWeatherClient, mockWeatherResponseMapper, weatherProvider :=
		givenADefaultWeatherProviderWithMockedComponents()

	weatherAPIResponse := givenAWeatherAPIResponse()
	expectedWeatherForecast := givenAWeatherForecastContract()

	mockRestWeatherClient.On("GetForecastForCityRequest", testCityLatitude, testCityLongitude).
		Return(weatherAPIResponse, nil)
	mockWeatherResponseMapper.On("ToWeatherForecastContract", weatherAPIResponse).
		Return(expectedWeatherForecast)

	actualWeatherForecast, err := weatherProvider.GetForecastForCity(testCityLatitude, testCityLongitude)

	assert.NoError(t, err)
	assert.Equal(t, expectedWeatherForecast, actualWeatherForecast)
	thenAssertMocksExpectations(t, mockConfig, mockHTTPClient, mockRestWeatherClient,
		mockWeatherResponseMapper)
}

func TestDefaultWeatherProvider_GetForecastForCity_whenRestClientFailsThenReturnsError(t *testing.T) {
	mockConfig, mockHTTPClient, mockRestWeatherClient, mockWeatherResponseMapper, weatherProvider :=
		givenADefaultWeatherProviderWithMockedComponents()

	weatherAPIResponse := givenAWeatherAPIResponse()
	restClientErr := errors.New("error in rest client")
	expectedErrorMessage := "error making request to weatherAPI"

	mockRestWeatherClient.On("GetForecastForCityRequest", testCityLatitude, testCityLongitude).
		Return(weatherAPIResponse, restClientErr)

	_, err := weatherProvider.GetForecastForCity(testCityLatitude, testCityLongitude)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), restClientErr.Error())
	assert.Contains(t, err.Error(), expectedErrorMessage)
	thenAssertMocksExpectations(t, mockConfig, mockHTTPClient, mockRestWeatherClient,
		mockWeatherResponseMapper)
}

func givenAWeatherAPIResponse() response.WeatherAPIResponse {
	return response.WeatherAPIResponse{
		Forecast: givenAResponseForecast(),
	}
}

func givenAResponseForecast() response.Forecast {
	return response.Forecast{
		Forecastdays: []response.Forecastday{{
			Day: response.Day{
				Condition: response.Condition{
					Text: testConditionText,
				},
			},
		}},
	}
}

func givenAWeatherForecastContract() contracts.WeatherForecast {
	return contracts.WeatherForecast{
		Forecastdays: []contracts.Forecastday{{
			Day: contracts.Day{
				Condition: contracts.Condition{
					Text: testConditionText,
				},
			},
		}},
	}
}

func givenADefaultWeatherProviderWithMockedComponents() (*MockWeatherProviderClientConfig, *MockHTTPClient,
	*MockRestWeatherClient, *MockWeatherResponseMapper, *DefaultWeatherProvider) {
	mockConfig := new(MockWeatherProviderClientConfig)
	mockHTTPClient := new(MockHTTPClient)
	mockRestWeatherClient := new(MockRestWeatherClient)
	mockWeatherResponseMapper := new(MockWeatherResponseMapper)

	weatherProvider := &DefaultWeatherProvider{
		restWeatherClient: mockRestWeatherClient,
		mapper:            mockWeatherResponseMapper,
	}

	return mockConfig, mockHTTPClient, mockRestWeatherClient, mockWeatherResponseMapper, weatherProvider
}

func givenADefaultWeatherProviderWithInjectedComponents() (*MockWeatherProviderClientConfig, *MockHTTPClient,
	*DefaultWeatherProvider) {
	mockConfig := new(MockWeatherProviderClientConfig)
	mockHTTPClient := new(MockHTTPClient)

	weatherProvider := &DefaultWeatherProvider{
		restWeatherClient: client.NewDefaultRestWeatherClient(mockConfig, mockHTTPClient),
		mapper:            mappers.NewDefaultWeatherResponseMapper(),
	}

	return mockConfig, mockHTTPClient, weatherProvider
}

func thenAssertMocksExpectations(t *testing.T, mockConfig *MockWeatherProviderClientConfig,
	mockHTTPClient *MockHTTPClient, mockRestWeatherClient *MockRestWeatherClient,
	mockWeatherResponseMapper *MockWeatherResponseMapper) {
	mockConfig.AssertExpectations(t)
	mockHTTPClient.AssertExpectations(t)
	mockRestWeatherClient.AssertExpectations(t)
	mockWeatherResponseMapper.AssertExpectations(t)
}
