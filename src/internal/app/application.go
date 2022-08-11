package app

import (
	"bufio"
	"fmt"
	"musement/src/internal/configs"
	"musement/src/internal/core/usecase"
	"musement/src/internal/infrastructure/entrypoints/cmd/getforecast"
	"musement/src/internal/infrastructure/providers/httpclient"
	"musement/src/internal/infrastructure/providers/musement"
	"musement/src/internal/infrastructure/providers/weather"
	"os"
)

func StartApp() error {
	// Config
	config := configs.NewConfig()
	httpClient := httpclient.NewHTTPClient()

	// Providers
	defaultMusementProvider := musement.NewDefaultMusementProvider(config, httpClient)
	defaultWeatherProvider := weather.NewDefaultWeatherProvider(config, httpClient)

	// UseCase
	defaultGetForecastForCitiesUseCase := usecase.NewDefaultGetForecastForCitiesUseCase(defaultMusementProvider,
		defaultWeatherProvider)

	// Handler
	defaultGetForecastHandler := getforecast.NewDefaultGetForecastHandler(defaultGetForecastForCitiesUseCase)

	buf := bufio.NewScanner(os.Stdin)
	fmt.Print("Press ENTER to get forecast for musement cities")
	err := defaultGetForecastHandler.HandleGetForecast()
	buf.Scan()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
