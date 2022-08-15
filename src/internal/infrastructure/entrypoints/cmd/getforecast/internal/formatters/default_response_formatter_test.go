package formatters

import (
	"musement/src/internal/core/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testCityName   = "Berlin"
	testCondition1 = "Sunny"
	testCondition2 = "Cloudy"
)

func TestNewDefaultResponseFormatter(t *testing.T) {
	expectedResponseFormatter := &defaultResponseFormatter{}

	actualResponseFormatter := NewDefaultResponseFormatter()

	assert.Equal(t, expectedResponseFormatter, actualResponseFormatter)
}

func TestDefaultResponseFormatter_FormatCityForcasted_success(t *testing.T) {
	responseFormatter := &defaultResponseFormatter{}

	cityForecasted := givenACityForecasted()
	expectedString := "Processed city Berlin | Sunny - Cloudy"

	actualString := responseFormatter.FormatCityForcasted(cityForecasted)

	assert.Equal(t, expectedString, actualString)
}

func givenACityForecasted() entities.CityForecasted {
	return entities.CityForecasted{
		Name: testCityName,
		Forecasts: []entities.Forecast{{
			Order:       0,
			Description: testCondition1,
		}, {
			Order:       1,
			Description: testCondition2,
		}},
	}
}
