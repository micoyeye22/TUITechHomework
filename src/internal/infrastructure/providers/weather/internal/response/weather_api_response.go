package response

type WeatherAPIResponse struct {
	Forecast Forecast `json:"forecast"`
}

type Forecast struct {
	Forecastdays []Forecastday `json:"forecastday"`
}

type Forecastday struct {
	Day Day `json:"day"`
}

type Day struct {
	Condition Condition `json:"condition"`
}

type Condition struct {
	Text string `json:"text"`
}
