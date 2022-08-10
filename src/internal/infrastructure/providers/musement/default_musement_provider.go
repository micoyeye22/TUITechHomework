package musement

import (
	"github.com/pkg/errors"
	"musement/src/internal/core/contracts"
	"musement/src/internal/infrastructure/providers/musement/internal/client"
	"musement/src/internal/infrastructure/providers/musement/internal/mappers"
	"musement/src/internal/infrastructure/providers/musement/internal/response"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MusementProviderClientConfig interface {
	BaseURL() string
}

type restMusementClient interface {
	GetCitiesRequest() (response.MusementAPIResponse, error)
}

type musementResponseMapper interface {
	ToCityContract(responseCity response.City) contracts.City
}

type DefaultMusementProvider struct {
	restMusementClient restMusementClient
	mapper             musementResponseMapper
}

func NewDefaultMusementProvider(config MusementProviderClientConfig, httpClient HTTPClient) *DefaultMusementProvider {
	return &DefaultMusementProvider{
		restMusementClient: client.NewDefaultRestMusementClient(config, httpClient),
		mapper:             mappers.NewDefaultMusementResponseMapper(),
	}
}

func (p *DefaultMusementProvider) GetCities() ([]contracts.City, error) {
	musementAPIResponse, reqErr := p.restMusementClient.GetCitiesRequest()
	if reqErr != nil {
		return nil, errors.Wrap(reqErr, "error making request to musementAPI")
	}

	var cities []contracts.City
	for _, city := range musementAPIResponse.Cities {
		cityContract := p.mapper.ToCityContract(city)
		cities = append(cities, cityContract)
	}

	return cities, nil
}
