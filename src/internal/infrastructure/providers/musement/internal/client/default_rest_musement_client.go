package client

import (
	"encoding/json"
	"fmt"
	"io"
	"musement/src/internal/infrastructure/providers/musement/internal/response"
	"net/http"

	"github.com/pkg/errors"
)

type MusementProviderClientConfig interface {
	BaseURL() string
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type defaultRestMusementClient struct {
	musementProviderClientConfig MusementProviderClientConfig
	httpClient                   HTTPClient
}

func NewDefaultRestMusementClient(config MusementProviderClientConfig, httpClient HTTPClient) *defaultRestMusementClient {
	return &defaultRestMusementClient{
		musementProviderClientConfig: config,
		httpClient:                   httpClient,
	}
}

func (rc *defaultRestMusementClient) GetCitiesRequest() (response.MusementAPIResponse, error) {
	baseURL := rc.musementProviderClientConfig.BaseURL()

	url := fmt.Sprintf("%s/cities", baseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		apiError := errors.Wrapf(err, "error building request to get cities")
		return response.MusementAPIResponse{}, apiError
	}

	responses, err := rc.httpClient.Do(req)
	if err != nil {
		apiError := errors.Wrapf(err, "error doing request to get cities")
		return response.MusementAPIResponse{}, apiError
	}
	defer func() { _ = responses.Body.Close() }()

	if responses.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(responses.Body)
		apiError := errors.Errorf("invalid status in response getting cities code '%d' and body '%s'",
			responses.StatusCode, body)
		return response.MusementAPIResponse{}, apiError
	}

	var musementAPIResponse response.MusementAPIResponse
	decoder := json.NewDecoder(responses.Body)
	err = decoder.Decode(&musementAPIResponse)

	return musementAPIResponse, err
}
