package client

import (
	"encoding/json"
	"fmt"
	"musement/src/internal/infrastructure/providers/musement/config"
	"musement/src/internal/infrastructure/providers/musement/internal/response"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testBaseURL        = "host://api"
	testInvalidBaseURL = "$%&!invalidURL/!%$&"
	testCityID         = 123
	testCityName       = "test-city"
	testCityLatitude   = 24.5
	testCityLongitude  = 20.1
)

func TestNewDefaultRestMusementClient(t *testing.T) {
	httpClient, mockConfig, expectedRestMusementClient := givenADefaultRestMusementClientWithMockedComponents()

	actualRestMusementClient := NewDefaultRestMusementClient(mockConfig, httpClient)

	assert.Equal(t, expectedRestMusementClient, actualRestMusementClient)
	mockConfig.AssertExpectations(t)
}

func TestDefaultRestMusementClient_GetCitiesRequest_success(t *testing.T) {
	_, mockConfig, restMusementClient := givenADefaultRestMusementClientWithMockedComponents()
	musementRestClientConfig := givenAMusementRestClientConfig(testBaseURL)

	expectedMusementAPIResponse := givenAGetCitiesRequestResponse()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockConfig.On("MusementProviderClientConfig").Return(musementRestClientConfig)
	whenGetCitiesRequestReturnsResponse(http.StatusOK, expectedMusementAPIResponse)

	actualMusementAPIResponse, err := restMusementClient.GetCitiesRequest()

	assert.NoError(t, err)
	assert.Equal(t, expectedMusementAPIResponse, actualMusementAPIResponse)
	mockConfig.AssertExpectations(t)
}

func TestDefaultRestMusementClient_GetCitiesRequest_whenBuildingRequestFailsThenReturnsError(t *testing.T) {
	_, mockConfig, restMusementClient := givenADefaultRestMusementClientWithMockedComponents()
	musementRestClientConfig := givenAMusementRestClientConfig(testInvalidBaseURL)

	expectedErrorMessage := "error building request to get cities"

	mockConfig.On("MusementProviderClientConfig").Return(musementRestClientConfig)

	_, err := restMusementClient.GetCitiesRequest()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
	mockConfig.AssertExpectations(t)
}

func TestDefaultRestMusementClient_GetCitiesRequest_whenClientFailsThenReturnsError(t *testing.T) {
	_, mockConfig, restMusementClient := givenADefaultRestMusementClientWithMockedComponents()
	musementRestClientConfig := givenAMusementRestClientConfig(testBaseURL)

	expectedErrorMessage := "error doing request to get cities"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mockConfig.On("MusementProviderClientConfig").Return(musementRestClientConfig)

	_, err := restMusementClient.GetCitiesRequest()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
	mockConfig.AssertExpectations(t)
}

func TestDefaultRestMusementClient_GetCitiesRequest_whenResponseHasInvalidStatusThenReturnsError(t *testing.T) {
	_, mockConfig, restMusementClient := givenADefaultRestMusementClientWithMockedComponents()
	musementRestClientConfig := givenAMusementRestClientConfig(testBaseURL)

	httpStatus := http.StatusBadRequest
	responseStruct := response.MusementAPIResponse{}
	responseBody := givenAResponseBodyString(t, responseStruct)

	expectedErrorMessage := fmt.Sprintf("invalid status in response getting cities code '%d' and body '%s'",
		httpStatus, responseBody)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockConfig.On("MusementProviderClientConfig").Return(musementRestClientConfig)
	whenGetCitiesRequestReturnsResponse(httpStatus, responseStruct)

	_, err := restMusementClient.GetCitiesRequest()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
	mockConfig.AssertExpectations(t)
}

func givenAResponseBodyString(t *testing.T, musementAPIResponse response.MusementAPIResponse) string {
	responseBodyBytes, marshallErr := json.Marshal(musementAPIResponse)
	require.NoError(t, marshallErr)
	return string(responseBodyBytes)
}

func givenAMusementRestClientConfig(url string) config.MusementProviderClientConfig {
	return config.MusementProviderClientConfig{
		BaseURL: url,
	}
}

func whenGetCitiesRequestReturnsResponse(status int, musementAPIResponse response.MusementAPIResponse) {
	httpMockURL := fmt.Sprintf("%s/%s", testBaseURL, "cities")

	httpmock.RegisterResponder(
		http.MethodGet,
		httpMockURL,
		func(req *http.Request) (resp *http.Response, e error) {
			resp, err := httpmock.NewJsonResponse(status, musementAPIResponse)
			return resp, err
		})
}

func givenAGetCitiesRequestResponse() response.MusementAPIResponse {
	return response.MusementAPIResponse{
		Cities: []response.City{
			givenACityResponse(),
		}}
}

func givenACityResponse() response.City {
	return response.City{
		ID:        testCityID,
		Name:      testCityName,
		Latitude:  testCityLatitude,
		Longitude: testCityLongitude,
	}
}

func givenADefaultRestMusementClientWithMockedComponents() (*http.Client, *MockConfig, *defaultRestMusementClient) {
	httpClient := http.DefaultClient
	mockConfig := new(MockConfig)

	restMusementClient := &defaultRestMusementClient{
		config:     mockConfig,
		httpClient: httpClient,
	}

	return httpClient, mockConfig, restMusementClient
}
