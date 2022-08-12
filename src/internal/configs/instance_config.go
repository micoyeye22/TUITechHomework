package configs

import "os"

type InstanceConfig struct {
	scope string
}

func NewInstanceConfig() *InstanceConfig {
	return &InstanceConfig{
		scope: os.Getenv("SCOPE"),
	}
}
