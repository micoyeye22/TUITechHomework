package response

type WeatherAPIResponse struct {
	Forecast Forecast
}

type Forecast struct {
	Forecastdays []Forecastday
}

type Forecastday struct {
	Day Day
}

type Day struct {
	Condition Condition
}

type Condition struct {
	Text string
}
