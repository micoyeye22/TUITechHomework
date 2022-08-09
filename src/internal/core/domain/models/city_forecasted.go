package models

type CityForecasted struct {
	Name      string
	Forecasts []Forecast
}

type Forecast struct {
	Order       int
	Description string
}
