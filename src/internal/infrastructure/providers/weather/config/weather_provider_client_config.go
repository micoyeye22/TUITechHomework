package config

type WeatherProviderClientConfig struct {
	BaseURL            string `yaml:"BaseURL"`
	WeatherClientToken string `yaml:"WeatherClientToken"`
	ForecastDays       int    `yaml:"ForecastDays"`
}
