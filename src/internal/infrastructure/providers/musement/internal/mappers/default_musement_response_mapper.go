package mappers

import (
	"musement/src/internal/core/contracts"
	"musement/src/internal/infrastructure/providers/musement/internal/response"
)

type defaultMusementResponseMapper struct {
}

func NewDefaultMusementResponseMapper() *defaultMusementResponseMapper {
	return &defaultMusementResponseMapper{}
}

func (m *defaultMusementResponseMapper) ToCityContract(responseCity response.City) contracts.City {
	return contracts.City{
		ID:        responseCity.ID,
		Name:      responseCity.Name,
		Latitude:  responseCity.Latitude,
		Longitude: responseCity.Longitude,
	}
}
