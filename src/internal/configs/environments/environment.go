package environments

import (
	musementProviderConfig "musement/src/internal/infrastructure/providers/musement/config"
	weatherProviderConfig "musement/src/internal/infrastructure/providers/weather/config"
)

type EnvironmentConfig struct {
	MusementClientConfig musementProviderConfig.MusementProviderClientConfig `yaml:"MusementClientConfig"`
	WeatherClientConfig  weatherProviderConfig.WeatherProviderClientConfig   `yaml:"WeatherClientConfig"`
}
