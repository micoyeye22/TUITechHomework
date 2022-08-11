package getforecast

import (
	"musement/src/internal/core/domain/models"
	"musement/src/internal/infrastructure/entrypoints/cmd/getforecast/internal/formatters"
	"musement/src/internal/infrastructure/entrypoints/cmd/getforecast/internal/printers"

	"github.com/pkg/errors"
)

type GetForecastForCitiesUseCase interface {
	GetForecastForCities() ([]models.CityForecasted, error)
}

type responseFormatter interface {
	FormatCityForcasted(cityForecasted models.CityForecasted) string
}

type responsePrinter interface {
	PrintCity(cityFormatted string)
}

type DefaultGetForecastHandler struct {
	GetForecastForCitiesUseCase GetForecastForCitiesUseCase
	formatter                   responseFormatter
	printer                     responsePrinter
}

func NewDefaultGetForecastHandler(useCase GetForecastForCitiesUseCase) *DefaultGetForecastHandler {
	return &DefaultGetForecastHandler{
		GetForecastForCitiesUseCase: useCase,
		formatter:                   formatters.NewDefaultResponseFormatter(),
		printer:                     printers.NewDefaultResponsePrinter(),
	}
}

func (h *DefaultGetForecastHandler) HandleGetForecast() error {
	citiesForecasted, useCaseErr := h.GetForecastForCitiesUseCase.GetForecastForCities()
	if useCaseErr != nil {
		return errors.Wrap(useCaseErr, "error in getForecast processing")
	}

	for _, cityForecasted := range citiesForecasted {
		cityFormatted := h.formatter.FormatCityForcasted(cityForecasted)
		h.printer.PrintCity(cityFormatted)
	}
	return nil
}
