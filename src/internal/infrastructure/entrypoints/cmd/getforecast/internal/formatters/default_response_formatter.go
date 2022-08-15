package formatters

import (
	"fmt"
	"musement/src/internal/core/domain/entities"
)

const responseFormat = "Processed city %s | %s - %s"

type defaultResponseFormatter struct {
}

func NewDefaultResponseFormatter() *defaultResponseFormatter {
	return &defaultResponseFormatter{}
}

func (f *defaultResponseFormatter) FormatCityForcasted(cityForecasted entities.CityForecasted) string {
	return fmt.Sprintf(responseFormat,
		cityForecasted.Name, cityForecasted.Forecasts[0].Description, cityForecasted.Forecasts[1].Description)
}
